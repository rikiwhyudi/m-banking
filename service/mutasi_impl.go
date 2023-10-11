package service

import (
	"context"
	"e-wallet/models"
	"e-wallet/repositories"
)

type transactionServiceImpl struct {
	transactionRepository repositories.TransactionRepository
}

func NewServiceTransactionImpl(transactionRepository repositories.TransactionRepository) TransactionService {
	return &transactionServiceImpl{transactionRepository}
}

func (s *transactionServiceImpl) GetTransactionService(ctx context.Context, accountNumber int) ([]models.Transaction, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		response, err := s.transactionRepository.GetTransactionRepository(ctx, accountNumber)

		if err != nil {
			return nil, err
		}

		return response, err
	}

}
