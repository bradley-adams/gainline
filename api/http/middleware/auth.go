package middleware

import (
	"context"
	"net/http"
	"net/url"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Auth(logger zerolog.Logger, domain, audience string) gin.HandlerFunc {
	issuerURL, err := url.Parse("https://" + domain + "/")
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse auth0 issuer url")
	}

	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{audience},
	)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to create jwt validator")
	}

	m := jwtmiddleware.New(jwtValidator.ValidateToken)

	return func(ctx *gin.Context) {
		encounteredError := true

		var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			encounteredError = false
			ctx.Request = r.WithContext(context.WithValue(r.Context(), "user", r.Context().Value(jwtmiddleware.ContextKey{})))
			ctx.Next()
		}

		m.CheckJWT(handler).ServeHTTP(ctx.Writer, ctx.Request)

		if encounteredError {
			ctx.Abort()
		}
	}
}
