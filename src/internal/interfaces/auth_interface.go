package interfaces

import (
	"context"

	"github.com/AgufSamudra/subscription/src/internal/models"
)

type AuthServiceInterface interface {
	RegisterService(ctx context.Context, email, password string) (models.RegisterResponse, error)
	LoginService(ctx context.Context, email, password string) (models.LoginResponse, error)
}

type AuthRepositoryInterface interface {
	RegisterRepository(ctx context.Context, email, password string) (models.RegisterResponse, error)
	LoginRepository(ctx context.Context, email, password string) (models.LoginResponse, error)
}
