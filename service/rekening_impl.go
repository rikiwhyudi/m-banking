package service

import (
	"context"
	"fmt"
	accNumberdto "m-banking/dto/rekening"
	"m-banking/internal"
	"m-banking/pkg/postgresql"
	"m-banking/pkg/rabbitmq"
	"m-banking/repositories"
	"sync"
)

type accountNumberServiceImpl struct {
	accountNumberRepository repositories.AccountNumberRepository
	customerRepository      repositories.CustomerRepository
	publisher               internal.MessageBroker
	wg                      sync.WaitGroup
}

func NewServiceAccountNumberImpl(accountNumberRepository repositories.AccountNumberRepository) AccountNumberService {
	return &accountNumberServiceImpl{accountNumberRepository, repositories.NewRepositoryCustomerImpl(postgresql.DB), internal.NewPublisherImpl(rabbitmq.RabbitMQChannel), sync.WaitGroup{}}
}

func (s *accountNumberServiceImpl) GetBalanceService(ctx context.Context, accountNumber int) (*accNumberdto.AccountNumberResponse, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		data, err := s.accountNumberRepository.GetBalanceRepository(ctx, accountNumber)
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
}

func (s *accountNumberServiceImpl) DepositService(ctx context.Context, account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		deposit, err := s.accountNumberRepository.GetBalanceRepository(ctx, account.AccountNumber)
		if err != nil {
			return nil, err
		}

		var publishError error
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			publishError = s.publisher.PublishMessage(deposit.ID, "D", account.Amount, "deposit")
		}()

		s.wg.Wait()

		if publishError != nil {
			return nil, publishError
		}

		// update balance
		deposit.Balance += account.Amount
		data, err := s.accountNumberRepository.DepositRepository(ctx, deposit)
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
}

func (s *accountNumberServiceImpl) CashoutService(ctx context.Context, account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		cashout, err := s.accountNumberRepository.GetBalanceRepository(ctx, account.AccountNumber)
		if err != nil {
			return nil, err
		}

		if cashout.Balance < account.Amount {
			return nil, fmt.Errorf("balance not enough: %v", cashout.Balance)
		}

		var publishError error
		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			publishError = s.publisher.PublishMessage(cashout.ID, "C", account.Amount, "cashout")
		}()

		s.wg.Wait()

		if publishError != nil {
			return nil, err
		}

		// update balance
		cashout.Balance -= account.Amount
		data, err := s.accountNumberRepository.CashoutRepository(ctx, cashout)
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

}

func (s *accountNumberServiceImpl) TransferService(ctx context.Context, transfer accNumberdto.TransferBalanceRequest) (*accNumberdto.TransferBalanceResponse, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if transfer.SenderAccountNumber == transfer.ReceiverAccountNumber {
			return nil, fmt.Errorf("invalid: can't transfer to account_number %v", transfer.ReceiverAccountNumber)
		}

		send, err := s.accountNumberRepository.GetBalanceRepository(ctx, transfer.SenderAccountNumber)
		if err != nil {
			return nil, fmt.Errorf("invalid: account_number not registered %v", transfer.SenderAccountNumber)
		}

		if send.Balance < transfer.Amount {
			return nil, fmt.Errorf("balance not enough: %v", send.Balance)
		}

		customer, err := s.customerRepository.GetCustomerRepository(ctx, send.CustomerID)
		if err != nil {
			return nil, err
		}

		receive, err := s.accountNumberRepository.GetBalanceRepository(ctx, transfer.ReceiverAccountNumber)
		if err != nil {
			return nil, fmt.Errorf("invalid: account_number not registered %v", transfer.ReceiverAccountNumber)
		}

		var publishError error
		accountNumbersID := []int{send.ID, receive.ID}
		for _, accountNumberID := range accountNumbersID {
			s.wg.Add(1)
			go func(accountNumberID int) {
				defer s.wg.Done()
				publishError = s.publisher.PublishMessage(accountNumberID, "T", transfer.Amount, "transfer")
			}(accountNumberID)
		}

		s.wg.Wait()

		if publishError != nil {
			return nil, err
		}

		send.Balance -= transfer.Amount
		receive.Balance += transfer.Amount
		data, err := s.accountNumberRepository.TransferRepository(ctx, send, receive)
		if err != nil {
			return nil, err
		}

		response := &accNumberdto.TransferBalanceResponse{
			Name:           customer.Name,
			SenderNumber:   data.AccountNumber,
			Balance:        data.Balance,
			ReceiverNumber: transfer.ReceiverAccountNumber,
			Amount:         transfer.Amount,
		}

		return response, err

	}

}
