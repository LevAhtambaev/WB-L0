package main

import (
	"LO/pkg/dsn"
	"LO/pkg/handlers"
	nats_client "LO/pkg/nats-client"
	"LO/pkg/recovery"
	"LO/pkg/repository"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

const driver = "postgres"

func main() {
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

	orderRepository := repository.NewOrderRepositoryImpl(db)
	cacheRepository := repository.NewCacheRepositoryImpl()
	paymentsRepository := repository.NewPaymentsRepositoryImpl(db)
	deliveryRepository := repository.NewDeliveryRepositoryImpl(db)
	itemRepository := repository.NewItemRepositoryImpl(db)

	templates := template.Must(template.ParseGlob("./templates/*"))

	orderHandler := handlers.NewOrderHandler(templates, logger, cacheRepository)

	recoverService := recovery.NewRecoverService(
		logger,
		orderRepository,
		cacheRepository,
		deliveryRepository,
		paymentsRepository,
		itemRepository,
	)

	err = recoverService.Recover()
	if err != nil {
		log.Errorf("Failed to recover cache from db: %s", err)
	}

	client := nats_client.NATSClient{
		Logger:             logger,
		OrderRepository:    orderRepository,
		CacheRepository:    cacheRepository,
		DeliveryRepository: deliveryRepository,
		PaymentRepository:  paymentsRepository,
		ItemRepository:     itemRepository,
	}
	go client.Subscribe()

	r := mux.NewRouter()

	r.HandleFunc("/api/orders", orderHandler.GetOrder).Methods("GET")

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
