package db

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DB Suite")
}

var _ = Describe("Open", func() {
	var (
		mockDB   *sql.DB
		mock     sqlmock.Sqlmock
		origOpen func(string, string) (*sql.DB, error)
	)

	BeforeEach(func() {
		var err error
		mockDB, mock, err = sqlmock.New(sqlmock.MonitorPingsOption(true))
		Expect(err).NotTo(HaveOccurred())

		origOpen = openDB
		openDB = func(driver, dsn string) (*sql.DB, error) {
			return mockDB, nil
		}
	})

	AfterEach(func() {
		openDB = origOpen
	})

	It("opens and pings the database successfully", func() {
		mock.ExpectPing()

		conn, err := Open("postgres://test")

		Expect(err).NotTo(HaveOccurred())
		Expect(conn).NotTo(BeNil())
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})

	It("closes the database and returns an error if ping fails", func() {
		mock.ExpectPing().WillReturnError(errors.New("ping failed"))
		mock.ExpectClose()

		conn, err := Open("postgres://test")

		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("ping db"))
		Expect(conn).To(BeNil())
		Expect(mock.ExpectationsWereMet()).To(Succeed())
	})
})
