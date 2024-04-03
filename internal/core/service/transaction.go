package service

import (
	"context"
	"encoding/csv"
	"errors"
	"stori/internal/adapter/email"
	"stori/internal/core/domain"
	"stori/internal/core/port"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	timeFormatMonthDay = "1/2"
)

type TransactionService struct {
	repository   port.TransactionRepository
	emailService *email.EmailService
}

func NewTransactionService(repository port.TransactionRepository, emailService *email.EmailService) *TransactionService {
	return &TransactionService{
		repository,
		emailService,
	}
}

func (ts *TransactionService) CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	return ts.repository.CreateTransaction(ctx, transaction)
}

func (ts *TransactionService) GetTransactionsByAccountID(ctx context.Context, accountID uuid.UUID) ([]domain.Transaction, error) {
	return ts.repository.GetTransactionsByAccountID(ctx, accountID)
}

func (ts *TransactionService) UploadFile(ctx context.Context, uploadFile *domain.UploadFile) error {
	if uploadFile.File.Header.Get("Content-Type") != "text/csv" {
		return errors.New("File is not a CSV file")
	}

	file, err := uploadFile.File.Open()
	if err != nil {
		return err
	}
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return err
	}

	transactions := make([]domain.Transaction, 0)

	for i, row := range data {
		if i == 0 {
			continue
		}

		date, err := time.Parse(timeFormatMonthDay, row[1])
		if err != nil {
			return err
		}

		credit := row[2][0] == byte('+')

		amount, err := decimal.NewFromString(row[2][1:])
		if err != nil {
			return err
		}

		transactions = append(transactions, domain.Transaction{
			Date:   date,
			Credit: credit,
			Amount: amount,
		})
	}

	_, err = ts.repository.CreateTransactions(ctx, uploadFile.AccountID, transactions)
	return err
}

func (ts *TransactionService) SendEmailSummary(ctx context.Context, accountID uuid.UUID) error {
	transactions, err := ts.repository.GetTransactionsByAccountID(ctx, accountID)
	if err != nil {
		return err
	}

	summary := &domain.Summary{}

	for i, transaction := range transactions {
		if i == 0 {
			summary.TransactionByMonth = make(map[time.Month]int)
		}
		if transaction.Credit {
			summary.Total = summary.Total.Add(transaction.Amount)
			summary.TotalAverageCredit = summary.TotalAverageCredit.Add(transaction.Amount)
		} else {
			summary.Total = summary.Total.Sub(transaction.Amount)
			summary.TotalAverageDebit = summary.TotalAverageDebit.Sub(transaction.Amount)
		}

		month := transaction.Date.Month()
		summary.TransactionByMonth[month]++

		if i == len(transactions)-1 {
			summary.TotalAverageCredit = summary.TotalAverageCredit.Div(decimal.NewFromInt(2))
			summary.TotalAverageDebit = summary.TotalAverageDebit.Div(decimal.NewFromInt(2))
		}
	}

	return ts.emailService.SendEmail(ctx, &domain.Email{
		AccountID: accountID,
		Summary:   summary,
	})
}
