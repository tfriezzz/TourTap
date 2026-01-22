-- +goose Up
CREATE TYPE booking_status AS ENUM (
  'pending',
  'confirmed',
  'cancelled'
);

CREATE TABLE bookings (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  date DATE NOT NULL,
  status booking_status NOT NULL DEFAULT 'pending',
  group_id UUID NOT NULL,

  FOREIGN KEY (group_id) REFERENCES groups(id)
);


-- +goose Down
DROP TABLE bookings;
