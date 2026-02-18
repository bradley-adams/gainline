//go:build integration
// +build integration

package db

import (
	"database/sql"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/stretchr/testify/require"
)

func TestVerifySchemaUpToDate_Integration(t *testing.T) {
	dsn := setupTestDB(t)

	t.Run("fails when no migrations applied", func(t *testing.T) {
		err := VerifySchemaUpToDate(dsn) // <-- pass DSN string
		require.Error(t, err)
	})

	t.Run("succeeds when schema is up to date", func(t *testing.T) {
		db, err := sql.Open("postgres", dsn)
		require.NoError(t, err)
		t.Cleanup(func() { require.NoError(t, db.Close()) })

		driver, err := postgres.WithInstance(db, &postgres.Config{})
		require.NoError(t, err)

		src, err := iofs.New(migrationFiles, "migrations")
		require.NoError(t, err)

		m, err := migrate.NewWithInstance("iofs", src, "postgres", driver)
		require.NoError(t, err)
		defer m.Close()

		require.NoError(t, m.Up())

		// Now verify schema using DSN string
		require.NoError(t, VerifySchemaUpToDate(dsn))
	})
}
