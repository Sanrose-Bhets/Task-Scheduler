-- +goose up
ALTER TABLE applications ADD COLUMN date_applied TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;

-- +goose down
ALTER TABLE applications DROP COLUMN date_applied;