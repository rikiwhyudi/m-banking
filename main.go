package main

import (
	"e-wallet/database"
	"e-wallet/pkg/postgresql"
	"e-wallet/pkg/rabbitmq"
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

	// initialize RabbitMQ connection
	rabbitmq.RabbitMqInit()

	//initialize RabbitMq consumer
	rabbitmq.RabbitMqConsumerInit()

	// initialize Mux Router connection
	r := mux.NewRouter()

	// run database migrations
	database.RunMigration()

	// initialize grouping routes
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT .env is required...")
	}

	// run server
	server := new(http.Server)
	server.Handler = r
	server.Addr = ":" + port
	fmt.Println("server runing on port " + port + "...")
	if err = server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
