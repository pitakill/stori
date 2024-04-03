package port

import (
	"context"
	"stori/internal/core/domain"

	"github.com/google/uuid"
)

type AccountRepository interface {
	CreateAccount(context.Context, *domain.Account) (*domain.Account, error)
	GetAccountByUserID(context.Context, uuid.UUID) (*domain.Account, error)
}

type AccountService interface {
	CreateAccount(context.Context, *domain.Account) (*domain.Account, error)
	GetAccountByUserID(context.Context, uuid.UUID) (*domain.Account, error)
}
