-- +goose Up
-- +goose StatementBegin
-- Add proper user_id column to journal (user has many journals)
ALTER TABLE journal ADD COLUMN IF NOT EXISTS user_id INTEGER REFERENCES usr(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE journal DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd
