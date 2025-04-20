-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS usr
(
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    email       VARCHAR,
    password    VARCHAR
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE usr;
-- +goose StatementEnd
