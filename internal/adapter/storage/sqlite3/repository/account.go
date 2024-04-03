package repository

import (
	"context"
	"stori/internal/core/domain"

	"github.com/google/uuid"
)

type AccountRepository struct {
	queries *Queries
}

func NewAccountRepository(queries *Queries) *AccountRepository {
	return &AccountRepository{
		queries,
	}
}

func (ar *AccountRepository) CreateAccount(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	if _, err := uuid.Parse(account.UserID.String()); err != nil {
		return nil, err
	}

	accountCreated, err := ar.queries.CreateAccount(context.Background(), CreateAccountParams{
		ID:     uuid.New(),
		UserID: account.UserID,
		Bank:   account.Bank,
		Number: account.Number,
	})
	if err != nil {
		return nil, err
	}

	account.ID = accountCreated.ID
	account.CreatedAt = accountCreated.CreatedAt.Time
	account.UpdatedAt = accountCreated.UpdatedAt.Time

	return account, nil
}

func (ar *AccountRepository) GetAccountByUserID(ctx context.Context, userID uuid.UUID) (*domain.Account, error) {
	if _, err := uuid.Parse(userID.String()); err != nil {
		return nil, err
	}

	account, err := ar.queries.GetAccountByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Account{
		ID:        account.ID,
		UserID:    account.UserID,
		Bank:      account.Bank,
		Number:    account.Number,
		CreatedAt: account.CreatedAt.Time,
		UpdatedAt: account.UpdatedAt.Time,
	}, nil
}
