package services

import (
	"context"
	"database/sql"

	"github.com/AgufSamudra/subscription/src/internal/interfaces"
	"github.com/AgufSamudra/subscription/src/internal/models"
	"github.com/AgufSamudra/subscription/src/internal/repository"
)

type AuthService struct {
	db         *sql.DB
	repository interfaces.AuthRepositoryInterface
}

func NewAuthService(db *sql.DB) (interfaces.AuthServiceInterface, error) {
	return &AuthService{
		db:         db,
		repository: repository.NewAuthRepository(db),
	}, nil
}

func (s *AuthService) RegisterService(ctx context.Context, email, password string) (models.RegisterResponse, error) {
	return s.repository.RegisterRepository(ctx, email, password)
}

func (s *AuthService) LoginService(ctx context.Context, email, password string) (models.LoginResponse, error) {
	return s.repository.LoginRepository(ctx, email, password)
}
