package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	CustomerRoutes(r)
	AccountNumberRoutes(r)
	TransactionRoutes(r)
}
