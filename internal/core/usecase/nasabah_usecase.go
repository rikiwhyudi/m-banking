package usecase

import (
	"context"
	"m-banking/internal/core/domain/dto"
	"m-banking/internal/core/domain/models"
	"m-banking/internal/core/ports"
	"time"
)

type customerUsecaseImpl struct {
	customerRepository ports.CustomerRepository
}

func NewCustomerUsecase(customerRepository ports.CustomerRepository) ports.CustomerUsecase {
	return &customerUsecaseImpl{customerRepository}
}

func (u *customerUsecaseImpl) RegisterCustomerUsecase(ctx context.Context, customer dto.CustomerRequest) (*dto.CustomerResponse, error) {

	createCustomer := models.Customer{
		// ID:          int(time.Now().Unix()),
		Name:        customer.Name,
		Nik:         customer.Nik,
		PhoneNumber: customer.PhoneNumber,
		UserID:      customer.UserID,
	}

	createAccountNumber := models.AccountNumber{
		AccountNumber: int(time.Now().UnixNano()),
		// CustomerID:    createCustomer.ID,
		Balance: 0.00,
	}

	data, err := u.customerRepository.RegisterCustomerRepository(ctx, createCustomer, createAccountNumber)

	if err != nil {
		return nil, err
	}

	response := &dto.CustomerResponse{
		ID:            data.ID,
		Name:          data.Name,
		Nik:           data.Nik,
		Email:         data.User.Email,
		PhoneNumber:   data.PhoneNumber,
		Status:        data.User.Role,
		AccountNumber: data.AccountNumber,
	}

	return response, err
}
