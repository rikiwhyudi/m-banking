package repositories

import (
	"e-wallet/models"

	"gorm.io/gorm"
)

type transactionRepositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryTransactionImpl(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{db}
}

func (r *transactionRepositoryImpl) GetTransactionRepository(accountNumber int) ([]models.Transaction, error) {
	var err error
	var transactions []models.Transaction

	// start db transaction
	tx := r.db.Begin()
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

	// commit db transaction
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return transactions, err
	}

	return transactions, err
}

func (r *transactionRepositoryImpl) CreateTransactionReposity(transaction models.Transaction) (models.Transaction, error) {
	var err error

	tx := r.db.Begin()
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
		return transaction, err
	}

	return transaction, err
}
