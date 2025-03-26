package database

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "host=localhost user=postgres password=windows dbname=network sslmode=disable"
)

var (
	db             *sql.DB
	ConnExistError = errors.New("db connection exists right now")
)

func InitDB() error {
	if db != nil {
		return ConnExistError
	}

	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	if err = db.PingContext(context.Background()); err != nil {
		return err
	}

	return nil
}
