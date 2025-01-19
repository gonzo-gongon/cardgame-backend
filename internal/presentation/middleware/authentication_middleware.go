package middleware

import (
	"net/http"
	"original-card-game-backend/internal/presentation/presenter"
	"strings"

	"github.com/gin-gonic/gin"
)

const authorizationHeaderName = "Authorization"

const bearerPrefixFirstLetterUpper = "Bearer "

const bearerPrefixFirstLetterLower = "bearer "

const tokenContextKey = "AuthorizationToken"

type TokenNotSuppliedError struct{}

func (e *TokenNotSuppliedError) Error() string {
	return "Authorization header is required"
}

// Refresh Tokenの時だけ使うmiddleware
type TokenRefreshMiddleware struct {
	authenticationPresenter *presenter.AuthenticationPresenter
}

func getTokenFromAuthorizationHeader(ctx *gin.Context) (string, error) {
	authorization := ctx.GetHeader(authorizationHeaderName)
	if authorization == "" {
		return "", &TokenNotSuppliedError{}
	}

	token := ""

	if strings.HasPrefix(authorization, bearerPrefixFirstLetterUpper) {
		token = strings.TrimPrefix(authorization, bearerPrefixFirstLetterUpper)
	}

	if strings.HasPrefix(authorization, bearerPrefixFirstLetterLower) {
		token = strings.TrimPrefix(authorization, bearerPrefixFirstLetterLower)
	}

	return token, nil
}

func setToken(ctx *gin.Context, token string) {
	ctx.Set(tokenContextKey, token)
}

func GetToken(ctx *gin.Context) string {
	return ctx.GetString(tokenContextKey)
}

func (m *TokenRefreshMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getTokenFromAuthorizationHeader(ctx)
		if err != nil {
			ctx.JSON(
				http.StatusUnauthorized,
				m.authenticationPresenter.Error(err),
			)
			ctx.Abort()

			return
		}

		setToken(ctx, token)
		ctx.Next()
	}
}

func NewTokenRefreshMiddleware(
	authenticationPresenter *presenter.AuthenticationPresenter,
) (*TokenRefreshMiddleware, error) {
	return &TokenRefreshMiddleware{
		authenticationPresenter: authenticationPresenter,
	}, nil
}

type AuthenticationMiddleware struct {
	authenticationPresenter *presenter.AuthenticationPresenter
}

func (m *AuthenticationMiddleware) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := getTokenFromAuthorizationHeader(ctx)
		if err == nil {
			setToken(ctx, token)
			ctx.Next()

			return
		}

		ctx.Next()
	}
}

func NewAuthenticationMiddleware(
	authenticationPresenter *presenter.AuthenticationPresenter,
) (*AuthenticationMiddleware, error) {
	return &AuthenticationMiddleware{
		authenticationPresenter: authenticationPresenter,
	}, nil
}
