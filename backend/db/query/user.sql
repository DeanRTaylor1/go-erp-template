-- name: CreateUser :one
INSERT INTO users (
  username,
  first_name,
  last_name,
  email,
  password,
  avatar,
  last_login,
  status,
  role
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = $1 LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users
LIMIT $2
OFFSET $1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: UpdateUser :exec
UPDATE users
SET
  username = $2,
  first_name = $3,
  last_name = $4,
  email = $5,
  password = $6,
  status = $7,
  role = $8,
  avatar = $9,
  updated_at = NOW()
WHERE id = $1;

-- name: UpdateLastLogin :exec
UPDATE users
SET
  last_login = NOW()
WHERE id = $1;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;