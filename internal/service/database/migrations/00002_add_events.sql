-- +goose Up
CREATE TABLE event
(
    id         SERIAL PRIMARY KEY,
    product_id BIGINT    NOT NULL,
    type       TEXT      NOT NULL,
    status     SMALLINT  NOT NULL,
    payload    JSONB     NOT NULL,
    updated    TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE event;