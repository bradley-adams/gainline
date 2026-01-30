package db

import (
	"database/sql"
	"errors"
	"fmt"
	"io/fs"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// VerifySchemaUpToDate verifies that the database schema matches the
// migrations embedded in this binary. It does NOT apply migrations.
func VerifySchemaUpToDate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("migrate driver: %w", err)
	}

	src, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		return fmt.Errorf("migrate source: %w", err)
	}

	m, err := migrate.NewWithInstance(
		"iofs", src,
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("migrate init: %w", err)
	}
	defer func() { _, _ = m.Close() }()

	dbVersion, dirty, err := dbMigrationState(m)
	if err != nil {
		return err
	}

	if dirty {
		return fmt.Errorf("database schema is dirty at version %d", dbVersion)
	}

	latestVersion, err := latestEmbeddedVersion(src)
	if err != nil {
		return err
	}

	if dbVersion != latestVersion {
		return fmt.Errorf(
			"database schema out of date: db=%d, expected=%d",
			dbVersion,
			latestVersion,
		)
	}

	return nil
}

func dbMigrationState(m *migrate.Migrate) (version uint, dirty bool, err error) {
	version, dirty, err = m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			return 0, false, nil
		}
		return 0, false, fmt.Errorf("migration version check failed: %w", err)
	}

	return version, dirty, nil
}

func latestEmbeddedVersion(src source.Driver) (uint, error) {
	latest, err := src.First()
	if err != nil {
		if err == migrate.ErrNilVersion {
			return 0, nil
		}
		return 0, fmt.Errorf("read migrations: %w", err)
	}

	for {
		next, err := src.Next(latest)
		if err != nil {
			if err == migrate.ErrNilVersion || errors.Is(err, fs.ErrNotExist) {
				break
			}
			return 0, fmt.Errorf("read migrations: %w", err)
		}
		latest = next
	}

	return latest, nil
}
