-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE url RENAME COLUMN short_url TO path;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE url RENAME COLUMN path TO short_url;
