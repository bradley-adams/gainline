//go:build integration
// +build integration

package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	ctx       context.Context
	container testcontainers.Container
	dsn       string
)

func setupDB(t *testing.T) {
	t.Helper()

	ctx = context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:16",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "gainline_test",
			"POSTGRES_USER":     "gainline",
			"POSTGRES_PASSWORD": "gainline",
		},
		WaitingFor: wait.
			ForSQL("5432/tcp", "postgres", func(host string, port nat.Port) string {
				return fmt.Sprintf(
					"postgres://gainline:gainline@%s:%s/gainline_test?sslmode=disable",
					host, port.Port(),
				)
			}).
			WithStartupTimeout(10 * time.Second),
	}

	var err error
	container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	host, err := container.Host(ctx)
	require.NoError(t, err)

	port, err := container.MappedPort(ctx, "5432")
	require.NoError(t, err)

	dsn = fmt.Sprintf(
		"postgres://gainline:gainline@%s:%s/gainline_test?sslmode=disable",
		host, port.Port(),
	)
}

func teardownDB(t *testing.T) {
	t.Helper()

	if container != nil {
		require.NoError(t, container.Terminate(ctx))
	}
}

func TestDBIntegration(t *testing.T) {
	setupDB(t)
	t.Cleanup(func() {
		teardownDB(t)
	})

	t.Run("connects to Postgres successfully", func(t *testing.T) {
		db, err := Open(dsn)
		require.NoError(t, err)
		require.NotNil(t, db)

		t.Cleanup(func() {
			require.NoError(t, db.Close())
		})

		require.NoError(t, db.Ping())
	})

	t.Run("fails with invalid DSN", func(t *testing.T) {
		db, err := Open("postgres://bad:bad@localhost:1/nope")
		require.Error(t, err)
		require.Nil(t, db)
	})
}
