-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments
(
    payments_uuid uuid PRIMARY KEY,
    transaction   VARCHAR(255),
    request_id    VARCHAR(255),
    currency      VARCHAR(3),
    provider      VARCHAR(255),
    amount        INTEGER,
    payment_dt    BIGINT,
    bank          VARCHAR(255),
    delivery_cost INTEGER,
    goods_total   INTEGER,
    custom_fee    INTEGER
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE payments;
-- +goose StatementEnd
