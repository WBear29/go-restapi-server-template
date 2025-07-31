
-- +migrate Up
ALTER TABLE samples ADD COLUMN age integer;

-- +migrate Down
ALTER TABLE samples DROP COLUMN age;
