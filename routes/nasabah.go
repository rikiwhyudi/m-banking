package routes

import (
	"e-wallet/handlers"
	"e-wallet/pkg/postgresql"
	"e-wallet/repositories"
	"e-wallet/service"

	"github.com/gorilla/mux"
)

func CustomerRoutes(r *mux.Router) {

	customerRepository := repositories.NewRepositoryCustomerImpl(postgresql.DB)
	customerService := service.NewServiceCustomerImpl(customerRepository)

	h := handlers.NewHanlderCustomerImpl(customerService)

	r.HandleFunc("/daftar", h.RegisterCustomerHandler).Methods("POST")
}
