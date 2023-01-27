// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: accounts.sql

package dbs

import (
	"context"
)

const accountGet = `-- name: AccountGet :one
SELECT id, number, name
FROM accounts
WHERE id = ?
`

func (q *Queries) AccountGet(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, accountGet, id)
	var i Account
	err := row.Scan(&i.ID, &i.Number, &i.Name)
	return i, err
}

const accountNew = `-- name: AccountNew :exec
INSERT INTO accounts (number, name)
VALUES (?, ?)
`

type AccountNewParams struct {
	Number int64
	Name   string
}

func (q *Queries) AccountNew(ctx context.Context, arg AccountNewParams) error {
	_, err := q.db.ExecContext(ctx, accountNew, arg.Number, arg.Name)
	return err
}