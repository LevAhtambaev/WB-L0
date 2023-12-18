package repository

import (
	"LO/pkg/models"
	"context"
)

type Repository interface {
	GetOrder(ctx context.Context, orderID int) (models.OrderJSON, error)
}
