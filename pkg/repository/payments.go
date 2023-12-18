package repository

import (
	"LO/pkg/models"
	"context"
	"database/sql"
	"github.com/google/uuid"
)

type PaymentsRepositoryImpl struct {
	DB *sql.DB
}

func NewPaymentsRepositoryImpl(DB *sql.DB) *PaymentsRepositoryImpl {
	return &PaymentsRepositoryImpl{DB: DB}
}

func (pr *PaymentsRepositoryImpl) AddPayment(ctx context.Context, payment models.Payment) (uuid.UUID, error) {
	payment.PaymentUUID = uuid.New()
	_, err := pr.DB.ExecContext(ctx, "INSERT INTO payments VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);", payment.PaymentUUID, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDT, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	if err != nil {
		return uuid.UUID{}, err
	}
	return payment.PaymentUUID, nil
}

func (pr *PaymentsRepositoryImpl) GetPaymentByID(ctx context.Context, uuid uuid.UUID) (models.Payment, error) {
	row := pr.DB.QueryRowContext(ctx, "SELECT * FROM payments WHERE payments_uuid=$1", uuid)
	var payment models.Payment
	err := row.Scan(&payment.PaymentUUID, &payment.Transaction, &payment.RequestID, &payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDT, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
	if err != nil {
		return models.Payment{}, err
	}
	return payment, nil
}
