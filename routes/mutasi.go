package routes

import (
	"m-banking/handlers"
	"m-banking/pkg/postgresql"
	"m-banking/repositories"
	"m-banking/service"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {

	transactionRepository := repositories.NewRepositoryTransactionImpl(postgresql.DB)
	transactionService := service.NewServiceTransactionImpl(transactionRepository)

	h := handlers.NewHandlerTransactionImpl(transactionService)

	r.HandleFunc("/mutasi/{id}", h.GetTransactionHandler).Methods("GET")

}
