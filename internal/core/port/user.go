package port

import (
	"context"
	"stori/internal/core/domain"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(context.Context, *domain.User) (*domain.User, error)
	GetUserByEmail(context.Context, string) (*domain.User, error)
	GetUserByAccountID(context.Context, uuid.UUID) (*domain.User, error)
}

type UserService interface {
	Register(context.Context, *domain.User) (*domain.User, error)
	GetUserByEmail(context.Context, string) (*domain.User, error)
}
