package main

import (
	"LO/pkg/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	clusterID = "test-cluster"
	clientID  = "server"
	url       = "nats://localhost:4222"
	layout    = "2006-01-02T15:04:05Z"
	subject   = "subject"
)

func main() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url))
	if err != nil {
		log.Errorf("Failed connect to NATS Streaming subsystem: %s", err)
		return
	}

	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			log.Errorf("Failed to close connection: %s", err)
		}
	}(sc)

	date, err := time.Parse(layout, time.Now().Format(layout))
	if err != nil {
		log.Errorf("Failed parse time at server: %s", err)
		return
	}

	orderUUID := uuid.New()

	order := models.OrderJSON{
		OrderUUID:   orderUUID,
		OrderUID:    "b563feb7b2b84b6test",
		TrackNumber: "WBILMTESTTRACK",
		Entry:       "WBIL",
		Delivery: models.Delivery{
			DeliveryUUID: uuid.New(),
			Name:         "TEST",
			Phone:        "+9720000000",
			Zip:          "2639809",
			City:         "Kiryat Mozkin",
			Address:      "Ploshad Mira 15",
			Region:       "Kraiot",
			Email:        "test@gmail.com",
		},
		Payment: models.Payment{
			PaymentUUID:  uuid.New(),
			Transaction:  "b563feb7b2b84b6test",
			RequestID:    "123",
			Currency:     "RUB",
			Provider:     "wbpay",
			Amount:       1000,
			PaymentDT:    1637907727,
			Bank:         "alpha",
			DeliveryCost: 1500,
			GoodsTotal:   320,
			CustomFee:    10,
		},
		Items: []models.Item{
			{
				ItemUUID:    uuid.New(),
				OrderUUID:   orderUUID,
				ChrtID:      9934930,
				TrackNumber: "WBILMTESTTRACK",
				Price:       453,
				Rid:         "ab4219087a764ae0btest",
				Name:        "Mascaras",
				Sale:        10,
				Size:        "M",
				TotalPrice:  317,
				NmID:        2389212,
				Brand:       "Vivienne Sabo",
				Status:      202,
			},
			{
				ItemUUID:    uuid.New(),
				OrderUUID:   orderUUID,
				ChrtID:      9934931,
				TrackNumber: "WBILMTESTTRACK",
				Price:       14999,
				Rid:         "dasfq23fsdfasdfasdf2ds",
				Name:        "Car",
				Sale:        10,
				Size:        "L",
				TotalPrice:  1234,
				NmID:        2389213,
				Brand:       "Toyota",
				Status:      202,
			},
		},
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "mtest",
		ShardKey:          "9",
		SmID:              99,
		DateCreated:       date,
		OofShard:          "l",
	}

	msg, err := json.Marshal(order)
	if err != nil {
		log.Errorf("Error at json Marshal: %s", err)
		return
	}
	err = sc.Publish(subject, msg)
	if err != nil {
		log.Errorf("Error at publish message: %s", err)
		return
	}
	log.Println("Message published")
	time.Sleep(10 * time.Second)
}
