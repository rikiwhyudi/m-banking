package repositories

import (
	"context"
	"e-wallet/models"
)

type TransactionRepository interface {
	GetTransactionRepository(ctx context.Context, accountNumber int) ([]models.Transaction, error)
	CreateTransactionReposity(transaction models.Transaction) (models.Transaction, error)
}
