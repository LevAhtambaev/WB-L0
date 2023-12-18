-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders
(
    id                 SERIAL PRIMARY KEY,
    order_uid          VARCHAR(255),
    track_number       VARCHAR(255),
    entry              VARCHAR(255),
    delivery_id        INTEGER REFERENCES delivery (id),
    payment_id         INTEGER REFERENCES payments (id),
    locale             VARCHAR(20),
    internal_signature VARCHAR(255),
    customer_id        VARCHAR(255),
    delivery_service   VARCHAR(255),
    shardkey           VARCHAR(255),
    sm_id              INTEGER,
    date_created       TIMESTAMP WITHOUT TIME ZONE,
    oof_shard          VARCHAR(255)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders;
-- +goose StatementEnd
