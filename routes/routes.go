package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	CustomerRoutes(r)
	AccountNumberRoutes(r)
	TransactionRoutes(r)
}
