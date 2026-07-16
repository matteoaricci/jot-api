-- +goose Up
-- +goose StatementBegin
-- Enforce one account per email. Partial index (WHERE deleted_at IS NULL) so a
-- soft-deleted user's email can be reused, and because gorm scopes lookups to
-- live rows. This supersedes the plain idx_usr_email, which we drop as redundant.
DROP INDEX IF EXISTS idx_usr_email;
CREATE UNIQUE INDEX IF NOT EXISTS idx_usr_email_unique ON usr (email) WHERE deleted_at IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_usr_email_unique;
CREATE INDEX IF NOT EXISTS idx_usr_email ON usr (email);
-- +goose StatementEnd
