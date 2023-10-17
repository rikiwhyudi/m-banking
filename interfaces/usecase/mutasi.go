package usecase

import (
	"context"
	"m-banking/models"
)

type TransactionUsecase interface {
	GetTransactionUsecase(ctx context.Context, accountNumber int) ([]models.Transaction, error)
}