package service

import (
	"context"
	customerdto "m-banking/dto/nasabah"
)

type CustomerService interface {
	RegisterCustomerService(ctx context.Context, customer customerdto.CustomerRequest) (*customerdto.CustomerResponse, error)
}
