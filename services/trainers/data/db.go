package data

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DbConfig struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
	MaxIdleTime  string
}

// The openDB() function returns a sql.DB connection pool.
func OpenDB(cfg DbConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)

	duration, err := time.ParseDuration(cfg.MaxIdleTime)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(duration)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
