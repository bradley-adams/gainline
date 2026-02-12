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

func setupTestDB(t *testing.T) string {
	t.Helper()

	ctx := context.Background()

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

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	host, err := container.Host(ctx)
	require.NoError(t, err)

	port, err := container.MappedPort(ctx, "5432")
	require.NoError(t, err)

	dsn := fmt.Sprintf(
		"postgres://gainline:gainline@%s:%s/gainline_test?sslmode=disable",
		host, port.Port(),
	)

	t.Cleanup(func() {
		require.NoError(t, container.Terminate(ctx))
	})

	return dsn
}
