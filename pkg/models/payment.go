package models

import "github.com/google/uuid"

type Payment struct {
	PaymentUUID  uuid.UUID `json:"payment_uuid"`
	Transaction  string    `json:"transaction"`
	RequestID    string    `json:"request_id"`
	Currency     string    `json:"currency"`
	Provider     string    `json:"provider"`
	Amount       int       `json:"amount"`
	PaymentDT    int64     `json:"paymentDT"`
	Bank         string    `json:"bank"`
	DeliveryCost int       `json:"delivery_cost"`
	GoodsTotal   int       `json:"goods_total"`
	CustomFee    int       `json:"custom_fee"`
}
