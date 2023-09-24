package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(url string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", url)

	if err != nil {
		return nil, err
	}

	return db, nil
}
