package service

import (
	"context"
	"stori/internal/core/domain"
	"stori/internal/core/port"
)

type UserService struct {
	repository port.UserRepository
}

func NewUserService(repository port.UserRepository) *UserService {
	return &UserService{
		repository,
	}
}

func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	return us.repository.CreateUser(ctx, user)
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	return us.repository.GetUserByEmail(ctx, email)
}
