package mysqldb

import (
	"balances/configs"
	"balances/pkg/logger"
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", configs.New().GetMySQLConnectionString())
	if err != nil {
		logger.Fatal(context.TODO(), "error on connect database", logger.Fields{"error": err.Error()})
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
