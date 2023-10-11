package repositories

import (
	"context"
	"e-wallet/models"

	"gorm.io/gorm"
)

func NewRepositoryCustomerImpl(db *gorm.DB) CustomerRepository {
	return &repository{db}
}

func (r *repository) RegisterCustomerRepository(ctx context.Context, customer models.Customer, accountNumber models.AccountNumber) (models.Customer, error) {
	var err error

	select {
	case <-ctx.Done():
		return customer, ctx.Err()
	default:
		tx := r.db.Begin()

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				return
			}
		}()

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
			tx.Rollback()
			return customer, err
		}

		return customer, err
	}
}
