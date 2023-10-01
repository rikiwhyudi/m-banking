package repositories

import (
	"e-wallet/models"

	"gorm.io/gorm"
)

type customerRepositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryCustomerImpl(db *gorm.DB) CustomerRepository {
	return &customerRepositoryImpl{db}
}

func (r *customerRepositoryImpl) RegisterCustomerRepository(customer models.Customer, accountNumber models.AccountNumber) (models.Customer, error) {
	var err error

	// start db transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return customer, tx.Error
	}

	err = tx.Create(&customer).Error
	if err != nil {
		tx.Rollback()
		return customer, err
	}

	// accountNumber.CustomerID = customer.ID

	err = tx.Create(&accountNumber).Error
	if err != nil {
		tx.Rollback()
		return customer, err
	}

	err = tx.Model(&customer).Preload("AccountNumber").First(&customer).Error
	if err != nil {
		tx.Rollback()
		return customer, err
	}

	// commit db transaction
	err = tx.Commit().Error
	if err != nil {
		return customer, err
	}

	return customer, err
}
