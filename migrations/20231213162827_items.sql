-- +goose Up
-- +goose StatementBegin
CREATE TABLE items
(
    item_uuid    uuid PRIMARY KEY,
    order_uuid   uuid REFERENCES orders (order_uuid),
    chrt_id      INTEGER,
    track_number VARCHAR(255),
    price        INTEGER,
    rid          VARCHAR(255),
    name         VARCHAR(255),
    sale         INTEGER,
    size         VARCHAR(255),
    total_price  INTEGER,
    nm_id        INTEGER,
    brand        VARCHAR(255),
    status       INTEGER
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
