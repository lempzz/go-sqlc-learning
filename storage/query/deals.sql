-- name: DealsCount :one
SELECT COUNT(id)
FROM deals;

-- name: BiggestDealsAccounts :many
SELECT COUNT(id) as count, account_id
FROM deals
GROUP BY account_id
ORDER BY count DESC;

-- name: NewDeal :exec
INSERT INTO deals (account_id, type, open_price, close_price, open_time, close_time, profit)
VALUES (?, ?, ?, ?, ?, ?, ?);