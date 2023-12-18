package recovery

import (
	"LO/pkg/models"
	"LO/pkg/repository"
	"context"
	"go.uber.org/zap"
)

type RecoverService struct {
	Logger             *zap.SugaredLogger
	OrderRepository    repository.OrderRepository
	CacheRepository    repository.CacheRepository
	DeliveryRepository repository.DeliveryRepository
	PaymentRepository  repository.PaymentRepository
	ItemRepository     repository.ItemRepository
}

func NewRecoverService(logger *zap.SugaredLogger, orderRepository repository.OrderRepository, cacheRepository repository.CacheRepository, deliveryRepository repository.DeliveryRepository, paymentRepository repository.PaymentRepository, itemRepository repository.ItemRepository) *RecoverService {
	return &RecoverService{Logger: logger, OrderRepository: orderRepository, CacheRepository: cacheRepository, DeliveryRepository: deliveryRepository, PaymentRepository: paymentRepository, ItemRepository: itemRepository}
}

func (rs *RecoverService) Recover() error {
	ctx := context.Background()
	orders, err := rs.OrderRepository.GetOrders(ctx)
	if err != nil {
		rs.Logger.Infof("Failed to receive orders from DB: %v", err)
		return err
	}

	var cacheOrder models.OrderJSON

	for _, order := range orders {
		cacheOrder.OrderUUID = order.OrderUUID
		cacheOrder.OrderUID = order.OrderUID
		cacheOrder.TrackNumber = order.TrackNumber
		cacheOrder.Entry = order.Entry
		cacheOrder.Locale = order.Locale
		cacheOrder.InternalSignature = order.InternalSignature
		cacheOrder.CustomerID = order.CustomerID
		cacheOrder.DeliveryService = order.DeliveryService
		cacheOrder.ShardKey = order.ShardKey
		cacheOrder.SmID = order.SmID
		cacheOrder.DateCreated = order.DateCreated
		cacheOrder.OofShard = order.OofShard

		delivery, err := rs.DeliveryRepository.GetDeliveryByID(ctx, order.DeliveryUUID)
		if err != nil {
			rs.Logger.Infof("Failed to receive delivery from DB: %v", err)
			return err
		}

		payment, err := rs.PaymentRepository.GetPaymentByID(ctx, order.PaymentUUID)
		if err != nil {
			rs.Logger.Infof("Failed to receive payment from DB: %v", err)
			return err
		}

		items, err := rs.ItemRepository.GetItemsByID(ctx, order.OrderUUID)
		if err != nil {
			rs.Logger.Infof("Failed to receive items from DB: %v", err)
			return err
		}

		cacheOrder.Delivery = delivery
		cacheOrder.Payment = payment
		cacheOrder.Items = items

		err = rs.CacheRepository.AddOrder(cacheOrder)
		if err != nil {
			rs.Logger.Infof("Failed to record order in cache: %v", err)
			return err
		}
	}
	return nil
}
