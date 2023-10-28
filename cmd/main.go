package main

import (
	"fmt"
	"m-banking/internal/infrastructure"
	"m-banking/internal/migrations"
	databaseInit "m-banking/pkg/database/sql"
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
	databaseInit.PostgreSQL()

	// initialize RabbitMQ connection
	rabbitmq.RabbitMqInit()

	//run RabbitMq consumer
	infrastructure.RabbitMqConsumer()

	// initialize Mux Router connection
	r := mux.NewRouter()

	// run database migrations
	migrations.RunMigration()

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
