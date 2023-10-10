package service

import (
	accNumberdto "e-wallet/dto/rekening"
	"e-wallet/internal"
	"e-wallet/pkg/rabbitmq"
	"e-wallet/repositories"
	"fmt"
	"sync"
)

type accountNumberServiceImpl struct {
	accountNumberRepository repositories.AccountNumberRepository
	publisher               internal.MessageBroker
}

func NewServiceAccountNumberImpl(accountNumberRepository repositories.AccountNumberRepository) AccountNumberService {
	return &accountNumberServiceImpl{accountNumberRepository, internal.NewPublisherImpl(rabbitmq.RabbitMQChannel)}
}

func (s *accountNumberServiceImpl) GetBalanceService(accountNumber int) (*accNumberdto.AccountNumberResponse, error) {

	data, err := s.accountNumberRepository.GetBalanceRepository(accountNumber)
	if err != nil {
		return nil, err
	}

	response := &accNumberdto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	return response, err
}

func (s *accountNumberServiceImpl) DepositService(account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	deposit, err := s.accountNumberRepository.GetBalanceRepository(account.AccountNumber)
	if err != nil {
		return nil, err
	}

	var publishError error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		publishError = s.publisher.PublishMessage(deposit.ID, "D", account.Amount, "deposit")
		if err != nil {
			fmt.Printf("error publishing message: %v\n", err)
		}
	}()

	wg.Wait()

	if publishError != nil {
		return nil, publishError
	}

	// update balance
	deposit.Balance += account.Amount
	data, err := s.accountNumberRepository.DepositRepository(deposit)
	if err != nil {
		return nil, err
	}

	response := &accNumberdto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	return response, err
}

func (s *accountNumberServiceImpl) CashoutService(account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	cashout, err := s.accountNumberRepository.GetBalanceRepository(account.AccountNumber)
	if err != nil {
		return nil, err
	}

	if cashout.Balance < account.Amount {
		return nil, fmt.Errorf("balance not enough: %v", cashout.Balance)
	}

	var publishError error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		publishError = s.publisher.PublishMessage(cashout.ID, "C", account.Amount, "cashout")
		if err != nil {
			fmt.Printf("error publishing message: %v\n", err)
		}
	}()

	wg.Wait()

	if publishError != nil {
		return nil, err
	}

	// update balance
	cashout.Balance -= account.Amount
	data, err := s.accountNumberRepository.CashoutRepository(cashout)
	if err != nil {
		return nil, err
	}

	response := &accNumberdto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	return response, err

}
