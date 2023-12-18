package nats_client

import (
	"LO/pkg/models"
	"LO/pkg/repository"
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"go.uber.org/zap"
	"time"
)

type NATSClient struct {
	Logger             *zap.SugaredLogger
	OrderRepository    repository.OrderRepository
	CacheRepository    repository.CacheRepository
	DeliveryRepository repository.DeliveryRepository
	PaymentRepository  repository.PaymentRepository
	ItemRepository     repository.ItemRepository
}

const (
	clusterID = "test-cluster"
	clientID  = "client"
	url       = "nats://localhost:4222"
	subject   = "subject"
)

func (nc *NATSClient) Subscribe() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url))
	if err != nil {
		nc.Logger.Infof("Failed connect to NATS Streaming subsystem: %v", err)
		return
	}

	nc.Logger.Info("Client connected succesfully")

	defer func(sc stan.Conn) {
		err := sc.Close()
		if err != nil {
			nc.Logger.Infof("Failed to close connection: %v", err)
		}
	}(sc)

	var orderJSON models.OrderJSON
	var order models.Order
	ctx := context.Background()

	_, err = sc.Subscribe(subject, func(msg *stan.Msg) {
		err = msg.Ack()
		if err != nil {
			nc.Logger.Infof("Error at recieving msg: %v", err)
			return
		}
		nc.Logger.Info("Received message")

		err = json.Unmarshal(msg.Data, &orderJSON)
		if err != nil {
			nc.Logger.Infof("Error at unmarshal to orderJSON: %v", err)
			return
		}

		nc.Logger.Info(orderJSON)

		deliveryUUID, err := nc.DeliveryRepository.AddDelivery(ctx, orderJSON.Delivery)
		if err != nil {
			nc.Logger.Infof("Error at add delivery to DB: %v", err)
			return
		}

		paymentUUID, err := nc.PaymentRepository.AddPayment(ctx, orderJSON.Payment)
		if err != nil {
			nc.Logger.Infof("Error at add payment to DB: %v", err)
			return
		}

		order.OrderUUID = orderJSON.OrderUUID
		order.OrderUID = orderJSON.OrderUID
		order.TrackNumber = orderJSON.TrackNumber
		order.Entry = orderJSON.Entry
		order.Locale = orderJSON.Locale
		order.InternalSignature = orderJSON.InternalSignature
		order.CustomerID = orderJSON.CustomerID
		order.DeliveryService = orderJSON.DeliveryService
		order.ShardKey = orderJSON.ShardKey
		order.SmID = orderJSON.SmID
		order.DateCreated = orderJSON.DateCreated
		order.OofShard = orderJSON.OofShard
		order.DeliveryUUID = deliveryUUID
		order.PaymentUUID = paymentUUID

		err = nc.OrderRepository.AddOrder(ctx, order)
		if err != nil {
			nc.Logger.Infof("Error at add order to DB: %v", err)
			return
		}

		err = nc.ItemRepository.AddItems(ctx, orderJSON.Items, orderJSON.OrderUUID)
		if err != nil {
			nc.Logger.Infof("Error at add items to DB: %v", err)
			return
		}

		err = nc.CacheRepository.AddOrder(orderJSON)
		if err != nil {
			nc.Logger.Infof("Error at add order to cache: %v", err)
			return
		}

	}, stan.SetManualAckMode())
	if err != nil {
		nc.Logger.Infof("Error at subcribe: %v", err)
		return
	}

	for {
		time.Sleep(1 * time.Second)
	}

}
