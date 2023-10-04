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
	accountNumberRepository repositories.AccountNumberRepository
	transactionRepository   repositories.TransactionRepository
}

func NewServiceAccountNumberImpl(accountNumberRepository repositories.AccountNumberRepository, transactionRepository repositories.TransactionRepository) AccountNumberService {
	return &accountNumberServiceImpl{accountNumberRepository, transactionRepository}
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

	return response, nil
}

func (s *accountNumberServiceImpl) DepositService(account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	ch, err := rabbitmq.GetRabbitMQChannel()
	if err != nil {
		return nil, err
	}

	defer ch.Close()

	deposit, err := s.accountNumberRepository.GetBalanceRepository(account.AccountNumber)
	if err != nil {
		return nil, err
	}

	queue, err := ch.QueueDeclare(
		"deposit",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	err = s.publishMessage(ch, deposit.ID, "D", account.Amount, queue.Name)
	if err != nil {
		fmt.Printf("failed to publish message: %v", err)
	}

	// update balance
	deposit.Balance += account.Amount
	data, err := s.accountNumberRepository.DepositRepository(deposit)
	if err != nil {
		return nil, err
	}

	done := make(chan bool)

	go func() {
		err = s.consumeMessage(ch, queue.Name, done)
		if err != nil {
			fmt.Printf("Error consuming message: %v\n", err)
		}
	}()

	<-done

	response := &accNumberdto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	return response, nil
}

func (s *accountNumberServiceImpl) CashoutService(account accNumberdto.AccountNumberRequest) (*accNumberdto.AccountNumberResponse, error) {

	ch, err := rabbitmq.GetRabbitMQChannel()
	if err != nil {
		return nil, err
	}

	defer ch.Close()

	cashout, err := s.accountNumberRepository.GetBalanceRepository(account.AccountNumber)
	if err != nil {
		return nil, err
	}

	if cashout.Balance < account.Amount {
		return nil, fmt.Errorf("balance not enough: %v", cashout.Balance)
	}

	queue, err := ch.QueueDeclare(
		"cashout",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	err = s.publishMessage(ch, cashout.ID, "C", account.Amount, queue.Name)
	if err != nil {
		fmt.Printf("failed to publish message: %v", err)
	}

	// update balance
	cashout.Balance -= account.Amount

	data, err := s.accountNumberRepository.CashoutRepository(cashout)

	if err != nil {
		return nil, err
	}

	done := make(chan bool)

	go func() {
		err = s.consumeMessage(ch, queue.Name, done)
		if err != nil {
			fmt.Printf("Error consuming message: %v\n", err)
		}
	}()

	<-done

	response := &accNumberdto.AccountNumberResponse{
		ID:            data.ID,
		AccountNumber: data.AccountNumber,
		Balance:       data.Balance,
	}

	return response, nil

}

// publisher
func (s *accountNumberServiceImpl) publishMessage(ch *amqp.Channel, accountNumber int, TransactionCode string, amount float64, queueName string) error {

	mutasi := transactiondto.TransactionRequest{
		AccountNumberID: accountNumber,
		TransactionCode: TransactionCode,
		Amount:          amount,
		Date:            time.Now(),
	}

	mutasiMessage, err := json.Marshal(mutasi)

	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json;",
			Body:        mutasiMessage,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

// consumer
func (s *accountNumberServiceImpl) consumeMessage(ch *amqp.Channel, queueName string, done chan bool) error {

	msgs, err := ch.Consume(
		queueName,
		"",
		false,
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

			if _, err := s.transactionRepository.CreateTransactionReposity(transaction); err != nil {
				fmt.Printf("failed save mutation to database: %s\n", err)
				return
			}

			fmt.Printf("Mutation message from RabbitMQ: %+v\n", mutasi)
			msg.Ack(false)

			done <- true

		}(msg)
	}

	return nil
}
