package ports

import (
	"context"
	"m-banking/internal/core/domain/dto"
	"m-banking/internal/core/domain/models"
)

type AuthRepository interface {
	RegisterAuthRepository(ctx context.Context, user models.User) (models.User, error)
	LoginAuthRepository(ctx context.Context, email string) (models.User, error)
}

type AuthUsecase interface {
	RegisterAuthUsecase(ctx context.Context, user dto.AuthRequest) (*dto.AuthRegisterResponse, error)
	LoginAuthUsecase(ctx context.Context, user dto.AuthRequest) (*dto.AuthLoginResponse, error)
}
