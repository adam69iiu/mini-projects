
-- name: CreateUser :one 
INSERT INTO users (id,username,email,password)
VALUES ($1,$2,$3,$4)
returning *;

-- name: UpdateUser :one
UPDATE users
set username = $1,
email = $2,
password = $3,
image_url = $4
WHERE id = $5
returning *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users 
WHERE username = $1
LIMIT 1;
