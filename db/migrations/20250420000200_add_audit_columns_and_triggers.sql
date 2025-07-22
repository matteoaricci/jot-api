-- +goose Up
-- +goose StatementBegin
-- audit columns for entry and journal
ALTER TABLE entry ADD COLUMN IF NOT EXISTS created_by VARCHAR;
ALTER TABLE entry ADD COLUMN IF NOT EXISTS updated_by VARCHAR;
ALTER TABLE journal ADD COLUMN IF NOT EXISTS created_by VARCHAR;
ALTER TABLE journal ADD COLUMN IF NOT EXISTS updated_by VARCHAR;

-- trigger to update journal.updated_at when an entry changes
CREATE OR REPLACE FUNCTION touch_journal_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE journal SET updated_at = NOW() WHERE id = COALESCE(NEW.journal_id, OLD.journal_id);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS entry_touch_journal_updated_at ON entry;
CREATE TRIGGER entry_touch_journal_updated_at
AFTER INSERT OR UPDATE OR DELETE ON entry
FOR EACH ROW EXECUTE FUNCTION touch_journal_updated_at();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS entry_touch_journal_updated_at ON entry;
DROP FUNCTION IF EXISTS touch_journal_updated_at;
ALTER TABLE entry DROP COLUMN IF EXISTS created_by;
ALTER TABLE entry DROP COLUMN IF EXISTS updated_by;
ALTER TABLE journal DROP COLUMN IF EXISTS created_by;
ALTER TABLE journal DROP COLUMN IF EXISTS updated_by;
-- +goose StatementEnd
