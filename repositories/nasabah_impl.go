package repositories

import (
	"context"
	"m-banking/models"

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
		tx := r.db.WithContext(ctx).Begin()

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

		err = tx.Create(&accountNumber).Error
		if err != nil {
			tx.Rollback()
			return customer, err
		}

		err = tx.Preload("AccountNumber").First(&customer).Error
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

func (r *repository) GetCustomerRepository(ctx context.Context, id int) (models.Customer, error) {
	var err error
	var customer models.Customer

	select {
	case <-ctx.Done():
		return customer, err
	default:
		tx := r.db.WithContext(ctx).Begin()

		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				return
			}
		}()

		if tx.Error != nil {
			return customer, tx.Error
		}

		err = tx.First(&customer, "id=?", id).Error
		if err != nil {
			return customer, err
		}

		err = tx.Commit().Error
		if err != nil {
			tx.Rollback()
			return customer, err
		}

		return customer, err
	}

}
