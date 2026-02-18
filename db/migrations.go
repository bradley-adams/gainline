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

	_ "github.com/lib/pq"
)

func VerifySchemaUpToDate(dbURL string) error {
	m, src, closeFn, err := newMigrator(dbURL)
	if err != nil {
		return err
	}
	defer closeFn()

	dbVersion, err := currentVersion(m)
	if err != nil {
		return err
	}

	latest, err := latestEmbeddedVersion(src)
	if err != nil {
		return err
	}

	if dbVersion != latest {
		return fmt.Errorf(
			"database schema out of date: db=%d expected=%d",
			dbVersion,
			latest,
		)
	}

	return nil
}

func newMigrator(dbURL string) (*migrate.Migrate, source.Driver, func(), error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, nil, nil, fmt.Errorf("ping db: %w", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		db.Close()
		return nil, nil, nil, fmt.Errorf("driver: %w", err)
	}

	src, err := iofs.New(migrationFiles, "migrations")
	if err != nil {
		db.Close()
		return nil, nil, nil, fmt.Errorf("source: %w", err)
	}

	m, err := migrate.NewWithInstance("iofs", src, "postgres", driver)
	if err != nil {
		db.Close()
		return nil, nil, nil, fmt.Errorf("migrate: %w", err)
	}

	closeFn := func() {
		m.Close()
		db.Close()
	}

	return m, src, closeFn, nil
}

func currentVersion(m *migrate.Migrate) (uint, error) {
	v, dirty, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			return 0, nil
		}
		return 0, fmt.Errorf("version: %w", err)
	}

	if dirty {
		return 0, fmt.Errorf("database schema is dirty at version %d", v)
	}

	return v, nil
}

func latestEmbeddedVersion(src source.Driver) (uint, error) {
	v, err := src.First()
	if err != nil {
		if err == migrate.ErrNilVersion {
			return 0, nil
		}
		return 0, err
	}

	for {
		next, err := src.Next(v)
		if err != nil {
			if err == migrate.ErrNilVersion || errors.Is(err, fs.ErrNotExist) {
				return v, nil
			}
			return 0, err
		}
		v = next
	}
}
