package repositories

import (
	"context"
	"e-wallet/models"
)

type CustomerRepository interface {
	RegisterCustomerRepository(ctx context.Context, customer models.Customer, accountNumber models.AccountNumber) (models.Customer, error)
}
