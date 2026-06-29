-- +goose Up
CREATE TABLE interviews (
    id SERIAL PRIMARY KEY,
    status VARCHAR(255) NOT NULL DEFAULT 'Not Performed',
    interviewDate DATE NOT NULL,
);

-- +goose Down
DROP TABLE interviews;