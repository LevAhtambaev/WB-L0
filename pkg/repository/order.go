package repository

import (
	"LO/pkg/models"
	"context"
	"database/sql"
)

type OrderRepositoryImpl struct {
	DB *sql.DB
}

func NewOrderRepositoryImpl(DB *sql.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{DB: DB}
}

func (or *OrderRepositoryImpl) GetOrders(ctx context.Context) ([]models.Order, error) {
	rows, err := or.DB.QueryContext(ctx, "SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	var orders []models.Order
	var order models.Order
	for rows.Next() {
		err = rows.Scan(&order.OrderUUID, &order.OrderUID, &order.TrackNumber, &order.Entry, &order.DeliveryUUID, &order.PaymentUUID, &order.Locale, &order.InternalSignature, &order.CustomerID, &order.DeliveryService, &order.ShardKey, &order.SmID, &order.DateCreated, &order.OofShard)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}

func (or *OrderRepositoryImpl) AddOrder(ctx context.Context, order models.Order) error {
	_, err := or.DB.ExecContext(ctx, "INSERT INTO orders VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14);", order.OrderUUID, order.OrderUID, order.TrackNumber, order.Entry, order.DeliveryUUID, order.PaymentUUID, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}
	return nil
}
