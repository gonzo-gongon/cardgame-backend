package gateway

import (
	"errors"
	"fmt"
	"original-card-game-backend/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SigningMethodError struct {
	alg interface{}
}

func (e *SigningMethodError) Error() string {
	return fmt.Sprintf("Unexpected signing method: %v", e.alg)
}

type InvalidTokenError struct{}

func (e *InvalidTokenError) Error() string {
	return "token is not valid"
}

type ParsingTokenError struct {
	cause error
}

func (e *ParsingTokenError) Error() string {
	return fmt.Sprintf("parsing token failed: %o", e.cause)
}

type AuthenticationConfig struct {
	secret []byte
}

type AuthenticationGateway interface {
	Generate(userID string) (string, error)
	GetUserID(tokenString string) (string, error)
	GetUserIDBypassTokenExpiry(tokenString string) (string, error)
	GetIssuedAt(tokenString string) (*time.Time, error)
}

type AuthenticationGatewayImpl struct {
	config AuthenticationConfig
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"userId"`
}

func (g *AuthenticationGatewayImpl) parseToken(
	tokenString string,
	claims jwt.Claims,
) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, &SigningMethodError{
					alg: token.Header["alg"],
				}
			}

			return g.config.secret, nil
		},
	)

	if err != nil {
		return nil, &ParsingTokenError{
			cause: err,
		}
	}

	return token, nil
}

func (g *AuthenticationGatewayImpl) verify(tokenString string) (*UserClaims, error) {
	claims := UserClaims{}

	token, err := g.parseToken(tokenString, &claims)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, &InvalidTokenError{}
	}

	return &claims, nil
}

func (g *AuthenticationGatewayImpl) verifyBypassExpiry(tokenString string) (*UserClaims, error) {
	claims := UserClaims{}

	_, err := g.parseToken(tokenString, &claims)
	if err != nil {
		if !errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}
	}

	return &claims, nil
}

func (g *AuthenticationGatewayImpl) Generate(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "system",
			Subject:   "user",
			Audience:  []string{"audience"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(100 * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
	})

	tokenString, err := token.SignedString(g.config.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (g *AuthenticationGatewayImpl) GetUserID(tokenString string) (string, error) {
	claims, err := g.verify(tokenString)
	if err != nil {
		return "", err
	}

	return claims.UserID, err
}

func (g *AuthenticationGatewayImpl) GetUserIDBypassTokenExpiry(tokenString string) (string, error) {
	claims, err := g.verifyBypassExpiry(tokenString)
	if err != nil {
		return "", err
	}

	return claims.UserID, err
}

func (g *AuthenticationGatewayImpl) GetIssuedAt(tokenString string) (*time.Time, error) {
	claims, err := g.verifyBypassExpiry(tokenString)
	if err != nil {
		return nil, err
	}

	issuedAt := claims.IssuedAt.Time

	return &issuedAt, nil
}

//nolint:ireturn // DIのためのコードなので許容する
func NewAuthenticationGateway(
	config *configs.Config,
) (AuthenticationGateway, error) {
	return &AuthenticationGatewayImpl{
		config: AuthenticationConfig{
			secret: []byte(config.JWT.Secret),
		},
	}, nil
}
