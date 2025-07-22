-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS usr
(
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP,
    email       VARCHAR,
    password    VARCHAR,
    first_name  VARCHAR,
    last_name   VARCHAR
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE usr;
-- +goose StatementEnd
