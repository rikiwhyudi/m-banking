package routes

import (
	"m-banking/internal/adapters/delivery/http"
	"m-banking/internal/adapters/repository"
	"m-banking/internal/core/usecase"
	"m-banking/pkg/database/sql"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {

	transactionRepository := repository.NewTransactionRepository(sql.DB)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository)

	h := http.NewTransactionHandler(transactionUsecase)

	r.HandleFunc("/mutasi/{id}", h.FindTransactionHandler).Methods("GET")

}
