package domain

import (
	"context"

	"github.com/roviery/catetin-api/entity"
	"github.com/roviery/catetin-api/models"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Store(ctx context.Context, user *models.User) (*models.User, error)
}

type UserUsecase interface {
	Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req entity.RegisterRequest) (*entity.RegisterResponse, error)
}
