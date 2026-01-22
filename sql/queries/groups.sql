-- name: CreateCustomer :one
INSERT INTO groups (id, created_at, updated_at, email, name, status)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1,
  $2,
  'unhandled'
  )
RETURNING *;


-- name: DeleteAllCustomers :exec
DELETE FROM groups;


-- name: GetCustomerByEmail :one
SELECT * FROM groups
WHERE email = $1;
