package usecase

import (
	"context"
	customerdto "m-banking/dto/nasabah"
)

type CustomerUsecase interface {
	RegisterCustomerUsecase(ctx context.Context, customer customerdto.CustomerRequest) (*customerdto.CustomerResponse, error)
}
