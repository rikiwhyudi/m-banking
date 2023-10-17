package usecase

import (
	"context"
	"m-banking/models"
)

type TransactionUsecase interface {
	FindTransactionUsecase(ctx context.Context, accountNumber int) ([]models.Transaction, error)
}
