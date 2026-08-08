package middleware

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"
)

var _ = Describe("Auth middleware", func() {
	var (
		router *gin.Engine
		logger zerolog.Logger
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		logger = zerolog.Nop()
		router = gin.New()
		router.Use(Auth(logger, "test.au.auth0.com", "https://api.dev.gainline.io"))
		router.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
		})
	})

	Describe("Auth", func() {
		It("should return 400 when no token is provided", func() {
			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)
			Expect(recorder.Code).To(Equal(http.StatusBadRequest))
		})

		It("should return 401 when an invalid token is provided", func() {
			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			req.Header.Set("Authorization", "Bearer invalid-token")
			recorder := httptest.NewRecorder()
			router.ServeHTTP(recorder, req)
			Expect(recorder.Code).To(Equal(http.StatusUnauthorized))
		})
	})
})
