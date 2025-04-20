-- +goose Up
-- +goose StatementBegin
ALTER TABLE journal ADD COLUMN IF NOT EXISTS usr_id INTEGER REFERENCES usr(id)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE journal DROP COLUMN user_id;
-- +goose StatementEnd