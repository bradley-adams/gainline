package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"

	mock_db "github.com/bradley-adams/gainline/db/db_handler/mock"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"go.uber.org/mock/gomock"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("health handlers", func() {
	var (
		ctrl   *gomock.Controller
		router *gin.Engine
		logger zerolog.Logger
		mockDB *mock_db.MockDB
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)

		ctrl = gomock.NewController(GinkgoT())
		mockDB = mock_db.NewMockDB(ctrl)

		logger = zerolog.Nop()

		router = gin.New()
		router.GET("/health", healthCheck(mockDB, logger))
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Describe("healthcheck", func() {
		It("returns 200 OK when HealthCheck returns nil", func() {
			mockDB.
				EXPECT().
				HealthCheck().
				Return(nil).
				Times(1)

			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Body.String()).To(ContainSubstring(`"message":"OK"`))
		})

		It("returns 503 when database health check fails", func() {
			mockDB.
				EXPECT().
				HealthCheck().
				Return(errors.New("db down")).
				Times(1)

			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusServiceUnavailable))
			Expect(w.Body.String()).To(ContainSubstring(`"status":"unhealthy"`))
		})
	})
})
