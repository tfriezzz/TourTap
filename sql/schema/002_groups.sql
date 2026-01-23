-- +goose Up
CREATE TYPE group_status AS ENUM (
'unhandled',
'payment_pending',
'confirmed',
'cancelled');


--TODO: handle 0 value for pax + existing id for tour_id

CREATE TABLE groups (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  email TEXT UNIQUE NOT NULL,
  name TEXT NOT NULL,
  pax INTEGER NOT NULL,
  status group_status NOT NULL DEFAULT 'unhandled',
  requested_tour_id INTEGER NOT NULL REFERENCES tours(id),
  requested_date DATE NOT NULL,
  booking_id INTEGER

);


-- +goose Down
DROP TABLE groups;
