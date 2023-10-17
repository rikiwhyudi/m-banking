package usecase

import (
	"context"
	models "m-banking/domain/models"
)

type TransactionUsecase interface {
	GetTransactionUsecase(ctx context.Context, accountNumber int) ([]models.Transaction, error)
}
