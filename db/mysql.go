package db

import (
	"github.com/chaiyapatoam/go-clean-fiber-boiler/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", config.GetEnv("MYSQL_URI"))

	if err != nil {
		return nil, err
	}

	return db, nil
}
