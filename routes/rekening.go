package routes

import (
	"m-banking/internal/adapters/delivery/http"
	"m-banking/internal/adapters/repository"
	"m-banking/internal/core/usecase"
	"m-banking/internal/infrastructure"
	"m-banking/pkg/database/sql"
	"m-banking/pkg/rabbitmq"

	"github.com/gorilla/mux"
)

func AccountNumberRoutes(r *mux.Router) {

	publisher := infrastructure.NewPublisher(rabbitmq.RabbitMQChannel)
	customerRepository := repository.NewCustomerRepository(sql.DB)
	accountNumberRepository := repository.NewAccountNumberRepository(sql.DB)
	accountNumberUsecase := usecase.NewAccountNumberUsecase(accountNumberRepository, customerRepository, publisher)

	h := http.NewAccountNumberHandler(accountNumberUsecase)

	r.HandleFunc("/saldo/{id}", h.GetBalanceHandler).Methods("GET")
	r.HandleFunc("/tabung", h.DepositHandler).Methods("PATCH")
	r.HandleFunc("/tarik", h.CashoutHandler).Methods("PATCH")
	r.HandleFunc("/transfer", h.TransferHandler).Methods("PATCH")
}
