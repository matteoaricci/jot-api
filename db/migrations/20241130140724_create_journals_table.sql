-- +goose Up
-- +goose StatementBegin
CREATE TABLE journal {
    id SERIAL PRIMARY KEY
    }
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
