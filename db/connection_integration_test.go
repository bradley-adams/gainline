//go:build integration
// +build integration

package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestDBIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Integration Suite")
}

var (
	ctx       context.Context
	container testcontainers.Container
	dsn       string
)

var _ = BeforeSuite(func() {
	ctx = context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:16",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "gainline_test",
			"POSTGRES_USER":     "gainline",
			"POSTGRES_PASSWORD": "gainline",
		},
		WaitingFor: wait.ForSQL("5432/tcp", "postgres", func(host string, port nat.Port) string {
			return fmt.Sprintf(
				"postgres://gainline:gainline@%s:%s/gainline_test?sslmode=disable",
				host, port.Port(),
			)
		}).WithStartupTimeout(5 * time.Second),
	}

	var err error
	container, err = testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	Expect(err).NotTo(HaveOccurred())

	host, err := container.Host(ctx)
	Expect(err).NotTo(HaveOccurred())

	port, err := container.MappedPort(ctx, "5432")
	Expect(err).NotTo(HaveOccurred())

	dsn = fmt.Sprintf(
		"postgres://gainline:gainline@%s:%s/gainline_test?sslmode=disable",
		host, port.Port(),
	)
})

var _ = AfterSuite(func() {
	if container != nil {
		_ = container.Terminate(ctx)
	}
})

var _ = Describe("Open (integration)", func() {
	It("connects to Postgres successfully", func() {
		db, err := Open(dsn)
		Expect(err).NotTo(HaveOccurred())
		Expect(db.Ping()).To(Succeed())
		Expect(db.Close()).To(Succeed())
	})

	It("fails with an invalid DSN", func() {
		db, err := Open("postgres://bad:bad@localhost:1/nope")
		Expect(err).To(HaveOccurred())
		Expect(db).To(BeNil())
	})
})
