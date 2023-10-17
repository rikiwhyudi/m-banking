package repository

import (
	"context"
	"m-banking/domain/models"
)

type CustomerRepository interface {
	RegisterCustomerRepository(ctx context.Context, customer models.Customer, accountNumber models.AccountNumber) (models.Customer, error)
	GetCustomerRepository(ctx context.Context, id int) (models.Customer, error)
}
