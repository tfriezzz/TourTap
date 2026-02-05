-- name: CreateBooking :one
INSERT INTO bookings (tour_id, created_at, updated_at, date)
  VALUES (
  $1,
  NOW(),
  NOW(),
  $2
  )
RETURNING *;


-- name: GetBookings :many
SELECT
  b.id AS booking_id,
  t.name AS tour_name,
  b.date,
  COUNT(g.id) AS group_count,
  COALESCE(SUM(g.pax), 0) AS total_pax,
  COALESCE(STRING_AGG(g.email, ', '), '') AS attending_groups
FROM bookings b
JOIN tours t ON b.tour_id = t.id
LEFT JOIN groups g ON g.booking_id = b.id
GROUP BY b.id, t.name, b.date
ORDER BY b.date DESC;

-- name: GetBookingByTourDate :one
SELECT * FROM bookings
WHERE date = $1 AND tour_id = $2;

-- name: GetAllBookingsOnDate :many
SELECT
  b.id AS booking_id,
  t.name AS tour_name,
  b.date,
  COUNT(g.id) AS group_count,
  COALESCE(SUM(g.pax), 0) AS total_pax,
  COALESCE(STRING_AGG(g.email, ', '), '') AS attending_groups
FROM bookings b
JOIN tours t ON b.tour_id = t.id
LEFT JOIN groups g ON g.booking_id = b.id
WHERE date = $1 AND NOT status = 'pending'
GROUP BY b.id, t.name, b.date
ORDER BY b.date DESC;
