-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE url (
  id SERIAL PRIMARY KEY,
  short_url VARCHAR(255) NOT NULL,
  long_url VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);
CREATE UNIQUE INDEX url_short_url_long_url_idx ON url (short_url) INCLUDE (long_url);


-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP INDEX url_short_url_long_url_idx;
DROP TABLE url;
