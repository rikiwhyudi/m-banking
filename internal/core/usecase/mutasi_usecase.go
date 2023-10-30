package usecase

import (
	"context"
	"m-banking/internal/core/domain/models"
	"m-banking/internal/core/ports"
)

type transactionUsecaseImpl struct {
	transactionRepository ports.TransactionRepository
}

func NewTransactionUsecase(transactionRepository ports.TransactionRepository) ports.TransactionUsecase {
	return &transactionUsecaseImpl{transactionRepository}
}

func (u *transactionUsecaseImpl) FindTransactionUsecase(ctx context.Context, accountNumber int) ([]models.Transaction, error) {

	response, err := u.transactionRepository.FindTransactionRepository(ctx, accountNumber)

	if err != nil {
		return nil, err
	}

	return response, err
}
