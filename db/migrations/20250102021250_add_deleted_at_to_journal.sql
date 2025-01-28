-- +goose Up
-- +goose StatementBegin
ALTER TABLE journal ADD COLUMN IF NOT EXISTS deleted_at TIMESTAMP
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
