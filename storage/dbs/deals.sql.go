// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: deals.sql

package dbs

import (
	"context"
	"time"
)

const biggestDealsAccounts = `-- name: BiggestDealsAccounts :many
SELECT COUNT(id) as count, account_id
FROM deals
GROUP BY account_id
ORDER BY count DESC
`

type BiggestDealsAccountsRow struct {
	Count     int64
	AccountID int64
}

func (q *Queries) BiggestDealsAccounts(ctx context.Context) ([]BiggestDealsAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, biggestDealsAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []BiggestDealsAccountsRow
	for rows.Next() {
		var i BiggestDealsAccountsRow
		if err := rows.Scan(&i.Count, &i.AccountID); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const dealsCount = `-- name: DealsCount :one
SELECT COUNT(id)
FROM deals
`

func (q *Queries) DealsCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, dealsCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const newDeal = `-- name: NewDeal :exec
INSERT INTO deals (account_id, type, open_price, close_price, open_time, close_time, profit)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type NewDealParams struct {
	AccountID  int64
	Type       int32
	OpenPrice  string
	ClosePrice string
	OpenTime   time.Time
	CloseTime  time.Time
	Profit     string
}

func (q *Queries) NewDeal(ctx context.Context, arg NewDealParams) error {
	_, err := q.db.ExecContext(ctx, newDeal,
		arg.AccountID,
		arg.Type,
		arg.OpenPrice,
		arg.ClosePrice,
		arg.OpenTime,
		arg.CloseTime,
		arg.Profit,
	)
	return err
}
