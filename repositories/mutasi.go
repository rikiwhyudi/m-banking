package repositories

import "e-wallet/models"

type TransactionRepository interface {
	GetTransactionRepository(accountNumber int) ([]models.Transaction, error)
	CreateTransactionReposity(transaction models.Transaction) (models.Transaction, error)
}
