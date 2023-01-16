-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
  id SERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(256), 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
)


-- +migrate StatementEnd
