-- +goose Up
-- +goose StatementBegin
ALTER TABLE journals ADD COLUMN deleted_at TIMESTAMP
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
