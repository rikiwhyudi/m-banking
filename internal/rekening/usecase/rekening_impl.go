package usecase

import (
	"context"
	"fmt"
	accNumberdto "m-banking/dto/rekening"
	"m-banking/interfaces/infrastructure/rabbitmq"
	"m-banking/interfaces/repository"
	"m-banking/interfaces/usecase"
	rabbitMq "m-banking/internal/infrastructure/rabbitmq"
	repositories "m-banking/internal/nasabah/repository"
	"m-banking/pkg/postgresql"
	ch "m-banking/pkg/rabbitmq"
	"sync"
)

type accountNumberUsecase struct {
	accountNumberRepository repository.AccountNumberRepository
	customerRepository      repository.CustomerRepository
	publisher               rabbitmq.Publisher
	wg                      sync.WaitGroup
}

func NewAccountNumberUsecaseImpl(accountNumberRepository repository.AccountNumberRepository) usecase.AccountNumberUsecase {
	return &accountNumberUsecase{accountNumberRepository, repositories.NewCustomerRepositoryImpl(postgresql.DB), rabbitMq.NewPublisherImpl(ch.RabbitMQChannel), sync.WaitGroup{}}
}

func (u *accountNumberUsecase) GetBalanceUsecase(ctx context.Context, accountNumber int) (*accNumberdto.AccountNumberResponse, error) {

	data, err := u.accountNumberRepository.GetBalanceRepository(ctx, accountNumber)
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

func (u *accountNumberUsecase) DepositUsecase(ctx context.Context, account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	deposit, err := u.accountNumberRepository.GetBalanceRepository(ctx, account.AccountNumber)
	if err != nil {
		return nil, err
	}

	var publishError error
	u.wg.Add(1)
	go func() {
		defer u.wg.Done()
		publishError = u.publisher.PublisherMessage(ctx, deposit.ID, "D", account.Amount, "deposit")
	}()

	u.wg.Wait()

	if publishError != nil {
		return nil, publishError
	}

	// update balance
	deposit.Balance += account.Amount
	data, err := u.accountNumberRepository.DepositRepository(ctx, deposit)
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

func (u *accountNumberUsecase) CashoutUsecase(ctx context.Context, account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	cashout, err := u.accountNumberRepository.GetBalanceRepository(ctx, account.AccountNumber)
	if err != nil {
		return nil, err
	}

	if cashout.Balance < account.Amount {
		return nil, fmt.Errorf("balance not enough: %v", cashout.Balance)
	}

	var publishError error
	u.wg.Add(1)
	go func() {
		defer u.wg.Done()
		publishError = u.publisher.PublisherMessage(ctx, cashout.ID, "C", account.Amount, "cashout")
	}()

	u.wg.Wait()

	if publishError != nil {
		return nil, err
	}

	// update balance
	cashout.Balance -= account.Amount
	data, err := u.accountNumberRepository.CashoutRepository(ctx, cashout)
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

func (u *accountNumberUsecase) TransferUsecase(ctx context.Context, transfer accNumberdto.TransferBalanceRequest) (*accNumberdto.TransferBalanceResponse, error) {

	if transfer.SenderAccountNumber == transfer.ReceiverAccountNumber {
		return nil, fmt.Errorf("invalid: can't transfer to account_number %v", transfer.ReceiverAccountNumber)
	}

	send, err := u.accountNumberRepository.GetBalanceRepository(ctx, transfer.SenderAccountNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid: account_number not registered %v", transfer.SenderAccountNumber)
	}

	if send.Balance < transfer.Amount {
		return nil, fmt.Errorf("balance not enough: %v", send.Balance)
	}

	customer, err := u.customerRepository.GetCustomerRepository(ctx, send.CustomerID)
	if err != nil {
		return nil, err
	}

	receive, err := u.accountNumberRepository.GetBalanceRepository(ctx, transfer.ReceiverAccountNumber)
	if err != nil {
		return nil, fmt.Errorf("invalid: account_number not registered %v", transfer.ReceiverAccountNumber)
	}

	var publishError error
	accountNumbersID := []int{send.ID, receive.ID}
	for _, accountNumberID := range accountNumbersID {
		u.wg.Add(1)
		go func(accountNumberID int) {
			defer u.wg.Done()
			publishError = u.publisher.PublisherMessage(ctx, accountNumberID, "T", transfer.Amount, "transfer")
		}(accountNumberID)
	}

	u.wg.Wait()

	if publishError != nil {
		return nil, err
	}

	send.Balance -= transfer.Amount
	receive.Balance += transfer.Amount
	data, err := u.accountNumberRepository.TransferRepository(ctx, send, receive)
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
