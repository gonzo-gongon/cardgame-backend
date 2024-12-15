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

func getTokenFromAuthorizationHeader(c *gin.Context) (string, error) {
	authorization := c.GetHeader(authorizationHeaderName)
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

func setToken(c *gin.Context, token string) {
	c.Set(tokenContextKey, token)
}

func GetToken(c *gin.Context) string {
	return c.GetString(tokenContextKey)
}

func (m *TokenRefreshMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getTokenFromAuthorizationHeader(c)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				m.authenticationPresenter.Error(err),
			)
			c.Abort()
			return
		}

		setToken(c, token)
		c.Next()
	}
}

func NewTokenRefreshMiddleware(
	authenticationPresenter *presenter.AuthenticationPresenter,
) (*TokenRefreshMiddleware, error) {
	return &TokenRefreshMiddleware{
		authenticationPresenter: authenticationPresenter,
	}, nil
}
