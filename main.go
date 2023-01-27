package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	acc, err := queries.AccountGet(context.Background(), 1); if err == nil {
		fmt.Printf("account: %v \n", acc)
	}

	queries.NewDeal(context.Background(), dbs.NewDealParams{
		AccountID: acc.ID,
		Type: 0,
		OpenPrice: "1.2",
		ClosePrice: "1.3",
		OpenTime: time.Date(2023, time.January, 25, 15, 0, 0, 0, time.UTC),
		CloseTime: time.Now().UTC(),
		Profit: "123.56",
	})
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