package repository

import (
	"LO/pkg/models"
	"context"
	"database/sql"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(DB *sql.DB) *OrderRepository {
	return &OrderRepository{DB: DB}
}

func (OrderRepository) GetOrder(ctx context.Context, orderID int) (models.OrderJSON, error) {
	return models.OrderJSON{}, nil
}
