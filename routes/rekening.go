package routes

import (
	"e-wallet/handlers"
	"e-wallet/internal"
	"e-wallet/pkg/postgresql"
	"e-wallet/pkg/rabbitmq"
	"e-wallet/repositories"
	"e-wallet/service"
	"fmt"

	"github.com/gorilla/mux"
)

func AccountNumberRoutes(r *mux.Router) {

	ch, err := rabbitmq.GetRabbitMqChannel()
	if err != nil {
		fmt.Println("failed to open channel: ", err)
	}

	publisher := internal.NewPublisherImpl(ch)
	accountNumberRepository := repositories.NewRepositoryAccountNumberImpl(postgresql.DB)
	accountNumberService := service.NewServiceAccountNumberImpl(accountNumberRepository, publisher)

	h := handlers.NewHandlerAccountNumberImpl(accountNumberService)

	r.HandleFunc("/saldo/{id}", h.GetBalanceHandler).Methods("GET")
	r.HandleFunc("/tabung", h.DepositHandler).Methods("PATCH")
	r.HandleFunc("/tarik", h.CashoutHandler).Methods("PATCH")
}
