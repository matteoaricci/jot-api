-- +goose Up
-- +goose StatementBegin
CREATE TABLE user
(
    id          SERIAL PRIMARY KEY,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    deleted_at TIMESTAMP,
    email VARCHAR,
    password VARCHAR
)

ALTER TABLE journal IF EXISTS ADD COLUMN user_id SERIAL REFERENCES user (id) ON DELETE CASCADE
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE user
ALTER TABLE journal IF EXISTS DROP COLUMN user_id
-- +goose StatementEnd
