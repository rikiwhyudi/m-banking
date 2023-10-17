package repository

import (
	"context"
	models "m-banking/domain/models"
)

type TransactionRepository interface {
	GetTransactionRepository(ctx context.Context, accountNumber int) ([]models.Transaction, error)
	CreateTransactionReposity(transaction models.Transaction) (models.Transaction, error)
}
