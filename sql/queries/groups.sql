-- name: CreateGroup :one
INSERT INTO groups (id, created_at, updated_at, email, name, pax, requested_tour_id, requested_date, booking_id)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
  )
RETURNING *;


-- name: DeleteAllGroups :exec
DELETE FROM groups;


-- name: GetGroupByEmail :one
SELECT * FROM groups
WHERE email = $1;


-- name: GroupStatusAccepted :one
UPDATE groups
SET status = 'accepted'
WHERE id = $1
RETURNING *;

