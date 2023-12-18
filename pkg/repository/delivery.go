package repository

import (
	"LO/pkg/models"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type DeliveryRepositoryImpl struct {
	DB *sql.DB
}

func NewDeliveryRepositoryImpl(DB *sql.DB) *DeliveryRepositoryImpl {
	return &DeliveryRepositoryImpl{DB: DB}
}

func (dr *DeliveryRepositoryImpl) AddDelivery(ctx context.Context, delivery models.Delivery) (uuid.UUID, error) {
	delivery.DeliveryUUID = uuid.New()
	_, err := dr.DB.ExecContext(ctx, "INSERT INTO delivery VALUES ($1,$2,$3,$4,$5,$6,$7,$8);", delivery.DeliveryUUID, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	if err != nil {
		return uuid.UUID{}, err
	}
	return delivery.DeliveryUUID, nil
}

func (dr *DeliveryRepositoryImpl) GetDeliveryByID(ctx context.Context, uuid uuid.UUID) (models.Delivery, error) {
	row := dr.DB.QueryRowContext(ctx, "SELECT * FROM delivery WHERE delivery_uuid=$1", uuid)
	var delivery models.Delivery
	err := row.Scan(&delivery.DeliveryUUID, &delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
	if err != nil {
		return models.Delivery{}, err
	}
	return delivery, nil
}
