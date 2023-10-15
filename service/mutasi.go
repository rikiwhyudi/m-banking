package service

import (
	"context"
	"m-banking/models"
)

type TransactionService interface {
	GetTransactionService(ctx context.Context, accountNumber int) ([]models.Transaction, error)
}
