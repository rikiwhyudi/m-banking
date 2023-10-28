package ports

import (
	"context"
	"m-banking/internal/core/domain/models"
)

type TransactionRepository interface {
	FindTransactionRepository(ctx context.Context, accountNumber int) ([]models.Transaction, error)
	CreateTransactionReposity(transaction models.Transaction) (models.Transaction, error)
}

type TransactionUsecase interface {
	FindTransactionUsecase(ctx context.Context, accountNumber int) ([]models.Transaction, error)
}
