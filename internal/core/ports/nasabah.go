package ports

import (
	"context"
	"m-banking/internal/core/domain/models"
	"m-banking/internal/dto"
)

type CustomerRepository interface {
	RegisterCustomerRepository(ctx context.Context, customer models.Customer, accountNumber models.AccountNumber) (models.Customer, error)
	GetCustomerRepository(ctx context.Context, id int) (models.Customer, error)
}

type CustomerUsecase interface {
	RegisterCustomerUsecase(ctx context.Context, customer dto.CustomerRequest) (*dto.CustomerResponse, error)
}
