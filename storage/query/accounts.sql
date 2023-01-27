-- name: AccountNew :exec
INSERT INTO accounts (number, name)
VALUES (?, ?);

-- name: AccountGet :one
SELECT *
FROM accounts
WHERE id = ?;