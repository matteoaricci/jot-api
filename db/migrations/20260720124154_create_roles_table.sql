-- +goose Up
-- +goose StatementBegin
CREATE TABLE role (
    id               SERIAL PRIMARY KEY,
    created_at       TIMESTAMP,
    updated_at       TIMESTAMP,
    deleted_at       TIMESTAMP,
    role_name        VARCHAR,
    role_description VARCHAR
);
-- +goose StatementEnd

-- +goose StatementBegin
INSERT INTO role (role_name, role_description) VALUES
    ('user', 'Standard user'),
    ('admin', 'Administrator with full access');
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE usr DROP COLUMN role;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE usr ADD COLUMN role_id INTEGER REFERENCES role(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE usr DROP COLUMN role_id;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE role;
-- +goose StatementEnd

-- +goose StatementBegin
ALTER TABLE usr ADD COLUMN role VARCHAR DEFAULT 'user';
-- +goose StatementEnd
