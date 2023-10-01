package main

import (
	"e-wallet/database"
	"e-wallet/pkg/postgresql"
	"e-wallet/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	// load configuration environment variables
	err := godotenv.Load()
	if err != nil {
		panic("failed to load godotenv")
	}

	// initialize DB connection
	postgresql.DatabaseInit()

	// initialize Mux Router connection
	r := mux.NewRouter()

	// run database migrations
	database.RunMigration()

	// initialize grouping routes
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	port := os.Getenv("PORT")

	// run server
	fmt.Println("server running on port " + port)
	http.ListenAndServe(":"+port, r)
}
