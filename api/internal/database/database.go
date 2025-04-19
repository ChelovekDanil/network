package database

import (
	"context"
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
)

const (
	driverName     = "postgres"
	dataSourceName = "host=localhost user=postgres dbname=network sslmode=disable"
)

var (
	db           *sql.DB
	ErrConnExist = errors.New("db connection exists right now")
)

func InitDB(ctx context.Context) error {
	if db != nil {
		return ErrConnExist
	}

	var err error
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}

	if err = db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}
