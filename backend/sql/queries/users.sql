-- name: CreateUser :one
INSERT INTO users (id, name, created_at, updated_at, email, hashed_password)
VALUES (
  gen_random_uuid(),
  $3,
  NOW(),
  NOW(),
  $1,
  $2
  )
RETURNING *;


-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;


-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1;


-- name: UpdateUserCredentials :one
UPDATE users
SET email = $2, hashed_password = $3, updated_at = NOW()
WHERE id = $1
RETURNING id, name, created_at, updated_at, email;
