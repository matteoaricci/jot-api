-- +goose Up
-- +goose StatementBegin
ALTER TABLE usr ADD COLUMN IF NOT EXISTS role VARCHAR DEFAULT 'user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE usr DROP COLUMN IF EXISTS role;
-- +goose StatementEnd
