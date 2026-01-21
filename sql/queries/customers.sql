-- name: CreateCustomer :one
INSERT INTO customers (id, created_at, updated_at, email)
VALUES (
  gen_random_uuid(),
  NOW(),
  NOW(),
  $1
  )
RETURNING *;


-- name: DeleteAllCustomers :exec
DELETE FROM customers;


-- name: GetCustomerByEmail :one
SELECT * FROM customers
WHERE email = $1;
