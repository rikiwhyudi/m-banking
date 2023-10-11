package service

import (
	"context"
	"e-wallet/models"
)

type TransactionService interface {
	GetTransactionService(ctx context.Context, accountNumber int) ([]models.Transaction, error)
}
