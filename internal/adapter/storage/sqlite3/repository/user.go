package repository

import (
	"context"
	"stori/internal/core/domain"

	"github.com/google/uuid"
)

type UserRepository struct {
	queries *Queries
}

func NewUserRepository(queries *Queries) *UserRepository {
	return &UserRepository{
		queries,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	userCreated, err := ur.queries.CreateUser(ctx, CreateUserParams{
		ID:        uuid.New(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})
	if err != nil {
		return nil, err
	}

	user.ID = userCreated.ID
	user.CreatedAt = userCreated.CreatedAt.Time
	user.UpdatedAt = userCreated.UpdatedAt.Time

	return user, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := ur.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}, nil
}

func (ur *UserRepository) GetUserByAccountID(ctx context.Context, accountID uuid.UUID) (*domain.User, error) {
	user, err := ur.queries.GetUserByAccountID(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: user.UpdatedAt.Time,
	}, nil
}
