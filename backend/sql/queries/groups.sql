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
SET status = 'accepted',
    updated_at = NOW()
WHERE id = $1 AND status = 'pending'
RETURNING *;

-- name: GroupStatusDeclined :one
UPDATE groups
SET status = 'declined',
    updated_at = NOW()
WHERE id = $1 AND status = 'pending'
RETURNING *;

-- name: GroupStatusConfirmed :one
UPDATE groups
SET status = 'confirmed',
    updated_at = NOW()
WHERE id = $1 AND status = 'accepted'
RETURNING *;
