-- +goose Up
-- +goose StatementBegin
CREATE TABLE items
(
    id           SERIAL PRIMARY KEY,
    order_id     INTEGER REFERENCES orders (id),
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
