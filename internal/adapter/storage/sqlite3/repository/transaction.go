package repository

import (
	"context"
	"stori/internal/core/domain"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionRepository struct {
	queries *Queries
}

func NewTransactionRepository(queries *Queries) *TransactionRepository {
	return &TransactionRepository{
		queries,
	}
}

func (tr *TransactionRepository) CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	if _, err := uuid.Parse(transaction.ID.String()); err != nil {
		return nil, err
	}

	transactionCreated, err := tr.queries.CreateTransaction(ctx, CreateTransactionParams{
		ID:        uuid.New(),
		AccountID: transaction.AccountID,
		Date:      transaction.Date,
		Credit:    transaction.Credit,
		Amount:    transaction.Amount.String(),
	})
	if err != nil {
		return nil, err
	}

	transaction.ID = transactionCreated.ID
	transaction.CreatedAt = transactionCreated.CreatedAt.Time
	transaction.UpdatedAt = transactionCreated.UpdatedAt.Time

	return transaction, nil
}

func (tr *TransactionRepository) GetTransactionsByAccountID(ctx context.Context, accountID uuid.UUID) ([]domain.Transaction, error) {
	if _, err := uuid.Parse(accountID.String()); err != nil {
		return nil, err
	}

	transactionsDB, err := tr.queries.GetTransactionsByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	transactions := make([]domain.Transaction, len(transactionsDB))

	for i, transaction := range transactionsDB {
		amount, _ := decimal.NewFromString(transaction.Amount)
		transactions[i].ID = transaction.ID
		transactions[i].AccountID = transaction.AccountID
		transactions[i].Date = transaction.Date
		transactions[i].Credit = transaction.Credit
		transactions[i].Amount = amount
		transactions[i].CreatedAt = transaction.CreatedAt.Time
		transactions[i].UpdatedAt = transaction.UpdatedAt.Time
	}

	return transactions, nil
}

func (tr *TransactionRepository) CreateTransactions(ctx context.Context, accountID uuid.UUID, transactions []domain.Transaction) ([]domain.Transaction, error) {
	account, err := tr.queries.GetAccountByID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	// sqlc does not have a method for bulk inserts in sqlite
	// https://github.com/sqlc-dev/sqlc/issues/3305
	for _, transaction := range transactions {
		transaction.AccountID = account.ID

		if _, err := tr.CreateTransaction(ctx, &transaction); err != nil {
			return nil, err
		}
	}

	return transactions, nil
}
