package service

import (
	transactiondto "e-wallet/dto/mutasi"
	accNumberdto "e-wallet/dto/rekening"
	"e-wallet/models"
	"e-wallet/pkg/rabbitmq"
	"e-wallet/repositories"
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

type accountNumberServiceImpl struct {
	accountNumberRepositoryImpl repositories.AccountNumberRepository
	transactionRepositoryImpl   repositories.TransactionRepository // tambahkan kontrak
}

func NewServiceAccountNumberImpl(accountNumberRepositoryImpl repositories.AccountNumberRepository, transactionRepositoryImpl repositories.TransactionRepository) AccountNumberService {
	return &accountNumberServiceImpl{accountNumberRepositoryImpl, transactionRepositoryImpl}
}

func (s *accountNumberServiceImpl) GetBalanceService(accountNumber int) (*accNumberdto.AccountNumberResponse, error) {

	data, err := s.accountNumberRepositoryImpl.GetBalanceRepository(accountNumber)
	if err != nil {
		return nil, err
	}

	response := &accNumberdto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	return response, nil
}

func (s *accountNumberServiceImpl) DepositService(account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	deposit, err := s.accountNumberRepositoryImpl.GetBalanceRepository(account.AccountNumber)
	if err != nil {
		return nil, err
	}

	ch, err := rabbitmq.GetRabbitMQChannel()
	if err != nil {
		return nil, err
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"mutasi_deposit",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	// create data mutation
	mutasi := transactiondto.TransactionRequest{
		AccountNumberID: deposit.ID,
		TransactionCode: "D",
		Amount:          account.Amount,
		Date:            time.Now(),
	}

	// publish to RabbitMQ
	mutasiMessage, err := json.Marshal(mutasi)
	if err != nil {
		return nil, err
	}

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json; charset=utf-8",
			Body:        mutasiMessage,
		},
	)

	if err != nil {
		return nil, err
	}

	// update balance
	deposit.Balance += account.Amount
	data, err := s.accountNumberRepositoryImpl.DepositRepository(deposit)
	if err != nil {
		return nil, err
	}

	response := &accNumberdto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	// consumer
	err = s.consumeMessage(queue.Name)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *accountNumberServiceImpl) CashoutService(account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	cashout, err := s.accountNumberRepositoryImpl.GetBalanceRepository(account.AccountNumber)
	if err != nil {
		return nil, err
	}

	if cashout.Balance < account.Amount {
		return nil, fmt.Errorf("balance not enough: %v", cashout.Balance)
	}

	ch, err := rabbitmq.GetRabbitMQChannel()
	if err != nil {
		return nil, err
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"mutasi_cashout",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	mutasi := transactiondto.TransactionRequest{
		AccountNumberID: cashout.ID,
		TransactionCode: "C",
		Amount:          account.Amount,
		Date:            time.Now(),
	}

	mutasiMessage, err := json.Marshal(mutasi)
	if err != nil {
		return nil, err
	}

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json; charset=utf-8",
			Body:        mutasiMessage,
		},
	)

	if err != nil {
		return nil, err
	}

	// update balance
	cashout.Balance -= account.Amount

	data, err := s.accountNumberRepositoryImpl.CashoutRepository(cashout)
	if err != nil {
		return nil, err
	}

	response := &accNumberdto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	err = s.consumeMessage(queue.Name)
	if err != nil {
		return nil, err
	}

	return response, nil

}

// func consumer msg from RabbitMQ
func (s *accountNumberServiceImpl) consumeMessage(queueName string) error {

	ch, err := rabbitmq.GetRabbitMQChannel()
	if err != nil {
		return err
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for msg := range msgs {
		go func(msg amqp.Delivery) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("Panic: %v\n", r)
					return
				}

			}()

			var mutasi transactiondto.TransactionRequest
			if err := json.Unmarshal(msg.Body, &mutasi); err != nil {
				fmt.Printf("failed parsing msg from RabbitMQ: %s\n", err)
				return
			}

			transaction := models.Transaction{
				AccountNumberID: mutasi.AccountNumberID,
				TransactionCode: mutasi.TransactionCode,
				Amount:          mutasi.Amount,
				Date:            mutasi.Date,
			}

			if _, err := s.transactionRepositoryImpl.CreateTransactionReposity(transaction); err != nil {
				fmt.Printf("failed save mutation to database: %s\n", err)
				return
			}

			fmt.Printf("Mutation message from RabbitMQ: %+v\n", mutasi)
			msg.Ack(false)

		}(msg)
	}

	return nil
}
