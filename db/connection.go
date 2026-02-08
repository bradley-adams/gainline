package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var openDB = sql.Open

// Open opens a database connection and verifies connectivity.
func Open(dsn string) (*sql.DB, error) {
	db, err := openDB("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}
