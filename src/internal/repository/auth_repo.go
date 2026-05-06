package repository

import (
	"context"
	"database/sql"

	"github.com/AgufSamudra/subscription/src/internal/interfaces"
	"github.com/AgufSamudra/subscription/src/internal/models"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) interfaces.AuthRepositoryInterface {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) RegisterRepository(ctx context.Context, email, password string) (models.RegisterResponse, error) {
	return models.RegisterResponse{
		Email: email,
	}, nil
}

func (r *AuthRepository) LoginRepository(ctx context.Context, email, password string) (models.LoginResponse, error) {
	return models.LoginResponse{}, nil
}
