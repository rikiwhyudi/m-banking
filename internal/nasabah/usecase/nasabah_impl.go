package usecase

import (
	"context"
	customerdto "m-banking/dto/nasabah"
	"m-banking/interfaces/repository"
	"m-banking/interfaces/usecase"
	"m-banking/models"
	"time"
)

type customerUsecase struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUsecaseImpl(customerRepository repository.CustomerRepository) usecase.CustomerUsecase {
	return &customerUsecase{customerRepository}
}

func (u *customerUsecase) RegisterCustomerUsecase(ctx context.Context, customer customerdto.CustomerRequest) (*customerdto.CustomerResponse, error) {

	createCustomer := models.Customer{
		ID:          int(time.Now().Unix()),
		Name:        customer.Name,
		Nik:         customer.Nik,
		PhoneNumber: customer.PhoneNumber,
	}

	createAccountNumber := models.AccountNumber{
		AccountNumber: int(time.Now().UnixNano()),
		CustomerID:    createCustomer.ID,
		Balance:       0.00,
	}

	data, err := u.customerRepository.RegisterCustomerRepository(ctx, createCustomer, createAccountNumber)

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
