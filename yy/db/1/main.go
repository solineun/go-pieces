package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"
)

const DRIVER_NAME = "postgres"

func main() {
	connStr := "user=yy-user password=yy-password dbname=yy-db sslmode=disable"

	db, err := sqlx.Open(DRIVER_NAME, connStr)
	if err != nil {
		slog.Error(fmt.Errorf("falied to get db conn: %w", err).Error())
		os.Exit(1)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		slog.Error(fmt.Errorf("falied to ping db: %w", err).Error())
		os.Exit(1)
	}
}
