-- name: CreateEntry :one
INSERT INTO payment.entries (
    account_id,
    amount
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM payment.entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM payment.entries
WHERE account_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;
