package port

import (
	"context"
	"stori/internal/core/domain"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	CreateTransaction(context.Context, *domain.Transaction) (*domain.Transaction, error)
	GetTransactionsByAccountID(context.Context, uuid.UUID) ([]domain.Transaction, error)
	CreateTransactions(context.Context, uuid.UUID, []domain.Transaction) ([]domain.Transaction, error)
}

type TransactionService interface {
	CreateTransaction(context.Context, *domain.Transaction) (*domain.Transaction, error)
	GetTransactionsByAccountID(context.Context, uuid.UUID) ([]domain.Transaction, error)
	UploadFile(context.Context, *domain.UploadFile) error
	SendEmailSummary(context.Context, uuid.UUID) error
}
