package repository

import (
	"context"

	"m-banking/models"
>>>>>>> clean-arch:interfaces/repository/mutasi.go
)

type TransactionRepository interface {
	GetTransactionRepository(ctx context.Context, accountNumber int) ([]models.Transaction, error)
	CreateTransactionReposity(transaction models.Transaction) (models.Transaction, error)
}
