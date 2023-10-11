package service

import (
	"context"
	customerdto "e-wallet/dto/nasabah"
)

type CustomerService interface {
	RegisterCustomerService(ctx context.Context, customer customerdto.CustomerRequest) (*customerdto.CustomerResponse, error)
}
