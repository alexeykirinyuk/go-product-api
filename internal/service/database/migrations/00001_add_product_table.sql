-- +goose Up
CREATE TABLE product
(
    id          SERIAL PRIMARY KEY,
    name        TEXT             NOT NULL,
    category    TEXT             NOT NULL,
    description TEXT             NOT NULL,
    brand       TEXT             NOT NULL,
    cost        DOUBLE PRECISION NOT NULL,
    currency    SMALLINT         NOT NULL,
    created     TIMESTAMP        NOT NULL,
    updated     TIMESTAMP        NOT NULL
);

-- +goose Down
DROP TABLE product;