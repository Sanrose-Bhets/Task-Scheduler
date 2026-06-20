-- +goose Up
CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
);
-- +goose Down
DROP TABLE accounts;


-- +goose Down
DROP TABLE accounts; 