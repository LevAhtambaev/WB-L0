package repository

import (
	"LO/pkg/models"
	"context"
	"github.com/google/uuid"
)

type CacheRepository interface {
	AddOrder(order models.OrderJSON) error
	GetOrderByID(orderUUID uuid.UUID) (models.OrderJSON, error)
}

type OrderRepository interface {
	AddOrder(ctx context.Context, order models.Order) error
	GetOrders(ctx context.Context) ([]models.Order, error)
}

type DeliveryRepository interface {
	AddDelivery(ctx context.Context, delivery models.Delivery) (uuid.UUID, error)
	GetDeliveryByID(ctx context.Context, uuid uuid.UUID) (models.Delivery, error)
}

type PaymentRepository interface {
	AddPayment(ctx context.Context, payment models.Payment) (uuid.UUID, error)
	GetPaymentByID(ctx context.Context, uuid uuid.UUID) (models.Payment, error)
}

type ItemRepository interface {
	AddItems(ctx context.Context, items []models.Item, orderUUID uuid.UUID) error
	GetItemsByID(ctx context.Context, orderUUID uuid.UUID) ([]models.Item, error)
}
