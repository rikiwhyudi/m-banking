package usecase

import (
	"context"
	"m-banking/interfaces/repository"
	"m-banking/interfaces/usecase"
	"m-banking/models"
)

type transactionUsecase struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionUsecaseImpl(transactionRepository repository.TransactionRepository) usecase.TransactionUsecase {
	return &transactionUsecase{transactionRepository}
}

func (s *transactionUsecase) GetTransactionUsecase(ctx context.Context, accountNumber int) ([]models.Transaction, error) {

	response, err := s.transactionRepository.GetTransactionRepository(ctx, accountNumber)

	if err != nil {
		return nil, err
	}

	return response, err
}