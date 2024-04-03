package service

import (
	"context"
	"stori/internal/core/domain"
	"stori/internal/core/port"

	"github.com/google/uuid"
)

type AccountService struct {
	repository port.AccountRepository
}

func NewAccountService(repository port.AccountRepository) *AccountService {
	return &AccountService{
		repository,
	}
}

func (as *AccountService) CreateAccount(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	return as.repository.CreateAccount(ctx, account)
}

func (as *AccountService) GetAccountByUserID(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	return as.repository.GetAccountByUserID(ctx, userID)
}
