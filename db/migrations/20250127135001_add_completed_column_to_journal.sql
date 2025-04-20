-- +goose Up
-- +goose StatementBegin
CREATE TYPE is_completed AS ENUM('true', 'false', 'unknown');
ALTER TABLE journal ADD COLUMN IF NOT EXISTS completed is_completed
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE journal DROP COLUMN completed;
DROP TYPE is_completed
-- +goose StatementEnd