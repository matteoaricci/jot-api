-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS journal
(
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    title       VARCHAR,
    description VARCHAR
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE journal;
-- +goose StatementEnd
