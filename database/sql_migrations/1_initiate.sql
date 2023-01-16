-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
  id SERIAL NOT NULL PRIMARY KEY,
  name VARCHAR(256), 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE books (
  id SERIAL NOT NULL PRIMARY KEY,
  title VARCHAR(256), 
  description VARCHAR(256), 
  image_url VARCHAR(256), 
  release_year INT,
  price VARCHAR(256), 
  total_page INT, 
  thickness VARCHAR(256), 
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  category_id INT
);

-- +migrate StatementEnd
