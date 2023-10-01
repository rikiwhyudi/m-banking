package service

import (
	customerdto "e-wallet/dto/nasabah"
	"e-wallet/models"
	"e-wallet/repositories"
	"time"
)

type customerServiceImpl struct {
	customerRepositoryImpl repositories.CustomerRepository
}

func NewServiceCustomerImpl(customerRepositoryImpl repositories.CustomerRepository) CustomerService {
	return &customerServiceImpl{customerRepositoryImpl}
}

func (s *customerServiceImpl) RegisterCustomerService(customer customerdto.CustomerRequest) (*customerdto.CustomerResponse, error) {

	createCustomer := models.Customer{
		ID:          int(time.Now().Unix()),
		Name:        customer.Name,
		Nik:         customer.Nik,
		PhoneNumber: customer.PhoneNumber,
	}

	createAccountNumber := models.AccountNumber{
		AccountNumber: int(time.Now().UnixNano()),
		CustomerID:    createCustomer.ID,
		Balance:       0,
	}

	data, err := s.customerRepositoryImpl.RegisterCustomerRepository(createCustomer, createAccountNumber)

	if err != nil {
		return nil, err
	}

	response := &customerdto.CustomerResponse{
		ID:            data.ID,
		Name:          data.Name,
		Nik:           data.Nik,
		PhoneNumber:   data.PhoneNumber,
		AccountNumber: data.AccountNumber,
	}

	return response, err
}
