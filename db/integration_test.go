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
	testCtx       context.Context
	testContainer testcontainers.Container
	testDSN       string
)

func setupTestDB(t *testing.T) string {
	t.Helper()

	testCtx = context.Background()

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
	testContainer, err = testcontainers.GenericContainer(testCtx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	host, err := testContainer.Host(testCtx)
	require.NoError(t, err)

	port, err := testContainer.MappedPort(testCtx, "5432")
	require.NoError(t, err)

	testDSN = fmt.Sprintf(
		"postgres://gainline:gainline@%s:%s/gainline_test?sslmode=disable",
		host, port.Port(),
	)

	t.Cleanup(func() {
		require.NoError(t, testContainer.Terminate(testCtx))
	})

	return testDSN
}
