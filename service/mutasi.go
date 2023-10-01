package service

import "e-wallet/models"

type TransactionService interface {
	GetTransactionService(accountNumber int) ([]models.Transaction, error)
}
