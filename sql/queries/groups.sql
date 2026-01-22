-- name: CreateCustomer :one
INSERT INTO groups (id, created_at, updated_at, email, name, pax, status, requested_tour_id, requested_date)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  $3,
  'unhandled',
  $4,
  $5
  )
RETURNING *;


-- name: DeleteAllCustomers :exec
DELETE FROM groups;


-- name: GetCustomerByEmail :one
SELECT * FROM groups
WHERE email = $1;
