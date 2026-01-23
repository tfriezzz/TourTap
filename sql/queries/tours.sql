-- NOTE: does id (serial) need input?

-- name: CreateTour :one
INSERT INTO tours (created_at, updated_at, name, base_price)
VALUES (
  NOW(),
  NOW(),
  $1,
  $2
  )
RETURNING *;

-- name: DeleteAllTours :exec
DELETE FROM tours;

-- name: GetTourByName :one
SELECT * FROM tours
WHERE name = $1;
