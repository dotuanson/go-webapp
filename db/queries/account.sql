-- name: CreateAccount :one
INSERT INTO payment.accounts (
    owner,
    balance,
    currency
) VALUES (
 $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM payment.accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM payment.accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM payment.accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE payment.accounts
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM payment.accounts
WHERE id = $1;