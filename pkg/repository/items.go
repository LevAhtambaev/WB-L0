package repository

import (
	"LO/pkg/models"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type ItemRepositoryImpl struct {
	DB *sql.DB
}

func NewItemRepositoryImpl(DB *sql.DB) *ItemRepositoryImpl {
	return &ItemRepositoryImpl{DB: DB}
}

func (ir *ItemRepositoryImpl) AddItems(ctx context.Context, items []models.Item, orderUUID uuid.UUID) error {
	for _, item := range items {
		_, err := ir.DB.ExecContext(ctx, "INSERT INTO items VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13);", item.ItemUUID, orderUUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ir *ItemRepositoryImpl) GetItemsByID(ctx context.Context, orderUUID uuid.UUID) ([]models.Item, error) {
	rows, err := ir.DB.QueryContext(ctx, "SELECT * FROM items WHERE order_uuid=$1", orderUUID)
	var items []models.Item
	var item models.Item
	for rows.Next() {
		err = rows.Scan(&item.ItemUUID, &item.OrderUUID, &item.ChrtID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err != nil {
		return nil, err
	}
	return items, nil
}
