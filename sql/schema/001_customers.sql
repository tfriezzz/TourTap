-- +goose Up
CREATE TYPE customer_status AS ENUM (
'unhandled',
'payment_pending',
'confirmed');

CREATE TABLE customers (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  email TEXT UNIQUE NOT NULL,
  name TEXT NOT NULL,
  status customer_status NOT NULL
);


-- +goose Down
DROP TABLE customers;
