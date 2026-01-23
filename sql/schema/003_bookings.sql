-- +goose Up
CREATE TYPE booking_status AS ENUM (
  'pending',
  'confirmed',
  'cancelled'
);

CREATE TABLE bookings (
  id SERIAL PRIMARY KEY,
  tour_id INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  date DATE NOT NULL,
  FOREIGN KEY (tour_id) REFERENCES tours(id),
  --status booking_status NOT NULL DEFAULT 'pending',
  CONSTRAINT unique_tour_date UNIQUE (tour_id, date)

);

ALTER TABLE groups ADD CONSTRAINT groups_booking_id FOREIGN KEY (booking_id) REFERENCES bookings(id);

-- +goose Down
DROP TABLE bookings;
