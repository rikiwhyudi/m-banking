package repository

import (
	"context"
	"m-banking/interfaces/repository"
	"m-banking/models"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepositoryImpl(db *gorm.DB) repository.TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetTransactionRepository(ctx context.Context, accountNumber int) ([]models.Transaction, error) {
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

func (r *transactionRepository) CreateTransactionReposity(transaction models.Transaction) (models.Transaction, error) {
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
