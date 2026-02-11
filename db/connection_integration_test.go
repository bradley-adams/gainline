//go:build integration
// +build integration

package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpen_Integration(t *testing.T) {
	dsn := setupTestDB(t)

	t.Run("connects successfully", func(t *testing.T) {
		db, err := Open(dsn)
		require.NoError(t, err)
		require.NotNil(t, db)
		require.NoError(t, db.Ping())
		require.NoError(t, db.Close())
	})

	t.Run("fails with invalid DSN", func(t *testing.T) {
		db, err := Open("postgres://bad:bad@localhost:1/nope")
		require.Error(t, err)
		require.Nil(t, db)
	})
}
