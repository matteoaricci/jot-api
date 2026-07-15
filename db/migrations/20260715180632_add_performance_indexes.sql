-- +goose Up
-- Index for finding journals by user (most common query pattern)
-- Used in: GET /api/users/:id/journals, user-specific journal listings
CREATE INDEX IF NOT EXISTS idx_journal_user_id ON journal(user_id);

-- Index for finding entries by journal (essential for journal detail views)
-- Used in: GET /api/journals/:id/entries
CREATE INDEX IF NOT EXISTS idx_entry_journal_id ON entry(journal_id);

-- Index for user lookup by email (used on every login)
-- Used in: POST /api/public/authenticate
CREATE INDEX IF NOT EXISTS idx_usr_email ON usr(email);

-- Index for filtering journals by completion status
-- Used in: GET /api/journals?completed=true/false
CREATE INDEX IF NOT EXISTS idx_journal_completed ON journal(completed);

-- Composite index for user's completed journals (optimizes common filtered query)
-- Used in: GET /api/users/:id/journals?completed=true
CREATE INDEX IF NOT EXISTS idx_journal_user_id_completed ON journal(user_id, completed);

-- +goose Down
DROP INDEX IF EXISTS idx_journal_user_id_completed;
DROP INDEX IF EXISTS idx_journal_completed;
DROP INDEX IF EXISTS idx_usr_email;
DROP INDEX IF EXISTS idx_entry_journal_id;
DROP INDEX IF EXISTS idx_journal_user_id;
