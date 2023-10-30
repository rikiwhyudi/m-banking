package usecase

import (
	"context"
	"fmt"
	"m-banking/internal/core/domain/models"
	"m-banking/internal/core/ports"
	"m-banking/internal/dto"
	"m-banking/pkg/bcrypt"
	jwtToken "m-banking/pkg/jwt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type authUsecaseImpl struct {
	authRepository ports.AuthRepository
}

func NewAuthUsecase(authRepository ports.AuthRepository) ports.AuthUsecase {
	return &authUsecaseImpl{authRepository}
}

func (u *authUsecaseImpl) RegisterAuthUsecase(ctx context.Context, user dto.AuthRequest) (*dto.AuthRegisterResponse, error) {

	password, err := bcrypt.HashingPassword(user.Password)
	if err != nil {
		return nil, err
	}

	createUser := models.User{
		Email:    user.Email,
		Password: password,
		Role:     "customer",
	}

	data, err := u.authRepository.RegisterAuthRepository(ctx, createUser)
	if err != nil {
		return nil, err
	}

	response := &dto.AuthRegisterResponse{
		ID:        data.ID,
		Email:     data.Email,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
	}

	return response, err
}

func (u *authUsecaseImpl) LoginAuthUsecase(ctx context.Context, user dto.AuthRequest) (*dto.AuthLoginResponse, error) {

	data, err := u.authRepository.LoginAuthRepository(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	isValid := bcrypt.CheckPassword(user.Password, data.Password)
	if !isValid {
		return nil, err
	}

	claims := jwt.MapClaims{}
	claims["id"] = data.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		fmt.Println(errGenerateToken)
		fmt.Println("unauthorized")
		return nil, errGenerateToken
	}

	response := &dto.AuthLoginResponse{
		ID:    data.ID,
		Email: data.Email,
		Role:  data.Role,
		Token: token,
	}

	return response, err
}
