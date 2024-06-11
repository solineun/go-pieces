package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func New(cfg Config) (*sqlx.DB, error) {
	dataSource := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", cfg.User, cfg.Password, cfg.Addr, cfg.Port, cfg.Db)

	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("faled to set db conn: %w", err)
	}

	return conn, nil
}
