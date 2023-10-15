package routes

import (
	"m-banking/handlers"
	"m-banking/pkg/postgresql"
	"m-banking/repositories"
	"m-banking/service"

	"github.com/gorilla/mux"
)

func CustomerRoutes(r *mux.Router) {

	customerRepository := repositories.NewRepositoryCustomerImpl(postgresql.DB)
	customerService := service.NewServiceCustomerImpl(customerRepository)

	h := handlers.NewHanlderCustomerImpl(customerService)

	r.HandleFunc("/daftar", h.RegisterCustomerHandler).Methods("POST")
}
