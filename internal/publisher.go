package internal

type Publisher interface {
	PublishMessage(accountNumber int, TransactionCode string, amount float64, queueName string) error
}
