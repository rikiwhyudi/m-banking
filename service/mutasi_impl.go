package service

import (
	"e-wallet/models"
	"e-wallet/repositories"
)

type transactionServiceImpl struct {
	transactionRepository repositories.TransactionRepository
}

func NewServiceTransactionImpl(transactionRepository repositories.TransactionRepository) TransactionService {
	return &transactionServiceImpl{transactionRepository}
}

func (s *transactionServiceImpl) GetTransactionService(accountNumber int) ([]models.Transaction, error) {

	response, err := s.transactionRepository.GetTransactionRepository(accountNumber)
	if err != nil {
		return nil, err
	}

	return response, err

}
