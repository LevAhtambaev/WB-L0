package main

import (
	"LO/pkg/dsn"
	"LO/pkg/handlers"
	"LO/pkg/repository"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

const driver = "postgres"

func main() {
	templates := template.Must(template.ParseGlob("./templates/*"))

	zapLogger, _ := zap.NewProduction()
	defer func(zapLogger *zap.Logger) {
		err := zapLogger.Sync()
		if err != nil {
			log.Errorf("Failed to sync zapLogger: %s", err)
		}
	}(zapLogger)
	logger := zapLogger.Sugar()

	err := godotenv.Load()
	if err != nil {
		log.Errorf("Failed to load env files: %s", err)
	}

	dbDSN := dsn.FromEnv()

	db, err := sql.Open(driver, dbDSN)
	if err != nil {
		log.Errorf("Failed to connect to db: %s", err)
	}

	sc, err := stan.Connect("test-cluster", "client-123")
	if err != nil {
		log.Errorf("Failed connetct to NATS: %s", err)
	}

	subscribtion, err := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.DurableName("my-durable"))

	if err != nil {
		log.Errorf("Failed to subscribe: %s", err)
	}

	orderRepository := repository.NewOrderRepository(db)
	orderHandler := handlers.NewOrderHandler(templates, logger, orderRepository)

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/order", orderHandler.GetOrder).Methods("GET")

	log.Println("server started")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		e := zapLogger.Sync()
		if e != nil {
			log.Printf("error at sync logger: %s", err)
		}
		log.Fatal(err)
	}
}
