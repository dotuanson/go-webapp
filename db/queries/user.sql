-- name: CreateUser :one
INSERT INTO payment.users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM payment.users
WHERE username = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM payment.users
WHERE username = $1;