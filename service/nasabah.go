package service

import customerdto "e-wallet/dto/nasabah"

type CustomerService interface {
	RegisterCustomerService(customer customerdto.CustomerRequest) (*customerdto.CustomerResponse, error)
}
