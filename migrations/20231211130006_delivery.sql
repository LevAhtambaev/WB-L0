-- +goose Up
-- +goose StatementBegin
CREATE TABLE delivery
(
    id      SERIAL PRIMARY KEY,
    name    VARCHAR(255),
    phone   VARCHAR(15),
    zip     VARCHAR(10),
    city    VARCHAR(255),
    address VARCHAR(255),
    region  VARCHAR(255),
    email   VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE delivery;
-- +goose StatementEnd
