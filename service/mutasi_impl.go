package service

import (
	"e-wallet/models"
	"e-wallet/repositories"
)

type transactionServiceImpl struct {
	transactionRepositoryImpl repositories.TransactionRepository
}

func NewServiceTransactionImpl(transactionRepositoryImpl repositories.TransactionRepository) TransactionService {
	return &transactionServiceImpl{transactionRepositoryImpl}
}

func (s *transactionServiceImpl) GetTransactionService(accountNumber int) ([]models.Transaction, error) {

	response, err := s.transactionRepositoryImpl.GetTransactionRepository(accountNumber)
	if err != nil {
		return nil, err
	}

	return response, err

}
