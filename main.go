package main

import (
	"context"
	"log"

	"github.com/lempzz/go-sqlc-learning/components"
	"github.com/lempzz/go-sqlc-learning/storage/dbs"
)

var queries *dbs.Queries
var logger *log.Logger

func main() {
	logger = log.Default()
	db := components.DBConnection()
	queries = dbs.New(db)
	
	newAccount("John Doe", 222333)
}

func newAccount(name string, num int64) {
	ctx := context.Background()

	err := queries.AccountNew(ctx, dbs.AccountNewParams{
		Name: name,
		Number: num,
	})

	if err != nil {
		logger.Printf("create account failed: %s", err.Error())
	}
}