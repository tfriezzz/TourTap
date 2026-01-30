-- +goose Up
CREATE TABLE tours (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  base_price DECIMAL(10, 2) NOT NULL

);


-- +goose Down
DROP TABLE tours;
