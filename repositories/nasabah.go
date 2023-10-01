package repositories

import "e-wallet/models"

type CustomerRepository interface {
	RegisterCustomerRepository(customer models.Customer, accountNumber models.AccountNumber) (models.Customer, error)
}
