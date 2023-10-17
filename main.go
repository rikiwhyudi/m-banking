package main

import (
	"fmt"
	"m-banking/database"
	receive "m-banking/internal/infrastructure/rabbitmq"
	"m-banking/pkg/postgresql"
	"m-banking/pkg/rabbitmq"
	"m-banking/routes"
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

	//run RabbitMq consumer
	receive.RabbitMqConsumer()

	// initialize Mux Router connection
	r := mux.NewRouter()

	// run database migrations
	database.RunMigration()

	// initialize and configure routes
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT .env is required...")
	}

	// run server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	fmt.Println("server runing on port " + port + "...")
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}

}
