package repository

import (
	"context"
	"m-banking/models"
)

type TransactionRepository interface {
	FindTransactionRepository(ctx context.Context, accountNumber int) ([]models.Transaction, error)
	CreateTransactionReposity(transaction models.Transaction) (models.Transaction, error)
}
