-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS entry
(
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    content     TEXT,
    journal_id  INTEGER REFERENCES journal(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE entry;
-- +goose StatementEnd
