package usecase

import (
	"context"
	"fmt"
	"m-banking/internal/core/ports"
	"m-banking/internal/dto"
	"m-banking/pkg/rabbitmq"
	"sync"
	"time"
)

type accountNumberUsecaseImpl struct {
	accountNumberRepository ports.AccountNumberRepository
	customerRepository      ports.CustomerRepository
	wg                      sync.WaitGroup
}

func NewAccountNumberUsecase(accountNumberRepository ports.AccountNumberRepository, customerRepository ports.CustomerRepository) ports.AccountNumberUsecase {
	return &accountNumberUsecaseImpl{accountNumberRepository, customerRepository, sync.WaitGroup{}}
}

func (u *accountNumberUsecaseImpl) GetBalanceUsecase(ctx context.Context, accountNumber int) (*dto.AccountNumberResponse, error) {

	data, err := u.accountNumberRepository.GetBalanceRepository(ctx, accountNumber)
	if err != nil {
		return nil, err
	}

	response := &dto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	return response, err
}

func (u *accountNumberUsecaseImpl) DepositUsecase(ctx context.Context, account dto.AccountNumberRequest) (*dto.AccountNumberResponse, error) {

	deposit, err := u.accountNumberRepository.GetBalanceRepository(ctx, account.AccountNumber)
	if err != nil {
		return nil, err
	}

	message := dto.TransactionRequest{
		AccountNumberID: deposit.ID,
		TransactionCode: "D",
		Amount:          account.Amount,
		Date:            time.Now(),
	}

	var publishError error
	u.wg.Add(1)
	go func() {
		defer u.wg.Done()
		publishError = rabbitmq.PublishMessage(ctx, "deposit", message)
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

	response := &dto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	return response, err
}

func (u *accountNumberUsecaseImpl) CashoutUsecase(ctx context.Context, account dto.AccountNumberRequest) (*dto.AccountNumberResponse, error) {

	cashout, err := u.accountNumberRepository.GetBalanceRepository(ctx, account.AccountNumber)
	if err != nil {
		return nil, err
	}

	if cashout.Balance < account.Amount {
		return nil, fmt.Errorf("balance not enough: %v", cashout.Balance)
	}

	message := dto.TransactionRequest{
		AccountNumberID: cashout.ID,
		TransactionCode: "C",
		Amount:          account.Amount,
		Date:            time.Now(),
	}

	var publishError error
	u.wg.Add(1)
	go func() {
		defer u.wg.Done()
		publishError = rabbitmq.PublishMessage(ctx, "cashout", message)
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

	response := &dto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	return response, err
}

func (u *accountNumberUsecaseImpl) TransferUsecase(ctx context.Context, transfer dto.TransferBalanceRequest) (*dto.TransferBalanceResponse, error) {

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

		message := dto.TransactionRequest{
			AccountNumberID: accountNumberID,
			TransactionCode: "T",
			Amount:          transfer.Amount,
			Date:            time.Now(),
		}

		u.wg.Add(1)
		go func(accountNumberID int) {
			defer u.wg.Done()
			publishError = rabbitmq.PublishMessage(ctx, "transfer", message)
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

	response := &dto.TransferBalanceResponse{
		Name:           customer.Name,
		SenderNumber:   data.AccountNumber,
		Balance:        data.Balance,
		ReceiverNumber: transfer.ReceiverAccountNumber,
		Amount:         transfer.Amount,
	}

	return response, err
}
