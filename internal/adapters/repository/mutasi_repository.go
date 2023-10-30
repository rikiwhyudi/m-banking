package repository

import (
	"context"
	"m-banking/internal/core/domain/models"
	"m-banking/internal/core/ports"

	"gorm.io/gorm"
)

func NewTransactionRepository(db *gorm.DB) ports.TransactionRepository {
	return &repository{db}
}

func (r *repository) FindTransactionRepository(ctx context.Context, accountNumber int) ([]models.Transaction, error) {
	var err error
	var transactions []models.Transaction

	tx := r.db.WithContext(ctx).Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return transactions, tx.Error
	}

	var account models.AccountNumber
	err = tx.First(&account, "account_number=?", accountNumber).Error
	if err != nil {
		tx.Rollback()
		return transactions, err
	}

	err = tx.Model(&account).Association("Transaction").Find(&transactions)
	if err != nil {
		tx.Rollback()
		return transactions, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return transactions, err
	}

	return transactions, err
}

func (r *repository) CreateTransactionReposity(transaction models.Transaction) (models.Transaction, error) {
	var err error

	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			return
		}
	}()

	if tx.Error != nil {
		return transaction, tx.Error
	}

	err = tx.Create(&transaction).Error
	if err != nil {
		tx.Rollback()
		return transaction, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return transaction, err
	}

	return transaction, err
}
