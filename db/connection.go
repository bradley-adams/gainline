package db

import (
	"database/sql"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

// Open opens a database connection and verifies connectivity.
func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}

// VerifySchemaUpToDate verifies that the database schema matches the
// migrations embedded in this binary. It does NOT apply migrations.
func VerifySchemaUpToDate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("migrate driver: %w", err)
	}

	sourceDriver, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("migrate source: %w", err)
	}

	m, err := migrate.NewWithInstance(
		"iofs", sourceDriver,
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("migrate init: %w", err)
	}

	dbVersion, dirty, err := m.Version()
	if err != nil && err != migrate.ErrNilVersion {
		return fmt.Errorf("migration version check failed: %w", err)
	}

	if dirty {
		return fmt.Errorf("database schema is dirty at version %d", dbVersion)
	}

	// Determine latest embedded migration version
	latest, err := sourceDriver.First()
	if err != nil {
		if err == migrate.ErrNilVersion {
			return nil
		}
		return fmt.Errorf("read migrations: %w", err)
	}

	for {
		next, err := sourceDriver.Next(latest)
		if err != nil {
			break
		}
		latest = next
	}

	if err == migrate.ErrNilVersion && latest != 0 {
		return fmt.Errorf(
			"database schema out of date: db=none, expected=%d",
			latest,
		)
	}

	if dbVersion != latest {
		return fmt.Errorf(
			"database schema out of date: db=%d, expected=%d",
			dbVersion,
			latest,
		)
	}

	return nil
}
