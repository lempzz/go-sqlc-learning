package components

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() *sql.DB {
	dsn := EnvString("GOOSE_DBSTRING")
	if dsn == "" {
		panic("does not except empty env GOOSE_DBSTRING")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	return db
}
