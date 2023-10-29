package routes

import (
	"m-banking/internal/adapters/delivery/http"
	"m-banking/internal/adapters/repository"
	"m-banking/internal/core/usecase"
	"m-banking/pkg/database/sql"

	"github.com/gorilla/mux"
)

func AccountNumberRoutes(r *mux.Router) {

	customerRepository := repository.NewCustomerRepository(sql.DB)
	accountNumberRepository := repository.NewAccountNumberRepository(sql.DB)
	accountNumberUsecase := usecase.NewAccountNumberUsecase(accountNumberRepository, customerRepository)

	h := http.NewAccountNumberHandler(accountNumberUsecase)

	r.HandleFunc("/saldo/{id}", h.GetBalanceHandler).Methods("GET")
	r.HandleFunc("/tabung", h.DepositHandler).Methods("PATCH")
	r.HandleFunc("/tarik", h.CashoutHandler).Methods("PATCH")
	r.HandleFunc("/transfer", h.TransferHandler).Methods("PATCH")
}
