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

type AuthenticationConfig struct {
	secret []byte
}

type AuthenticationGateway struct {
	config AuthenticationConfig
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserID string `json:"userId"`
}

func (g *AuthenticationGateway) parseToken(
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
	return token, err
}

func (g *AuthenticationGateway) verify(tokenString string) (*UserClaims, error) {
	claims := UserClaims{}

	token, err := g.parseToken(tokenString, &claims)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if !token.Valid {
		err := fmt.Errorf("token is not valid")

		fmt.Println(err)
		return nil, err
	}

	return &claims, nil
}

func (g *AuthenticationGateway) verifyBypassExpiry(tokenString string) (*UserClaims, error) {
	claims := UserClaims{}

	token, err := g.parseToken(tokenString, &claims)
	if err != nil {
		if !errors.Is(err, jwt.ErrTokenExpired) {
			fmt.Println(err)
			return nil, err
		}

		fmt.Printf("expired token: %o", token)
	}

	return &claims, nil
}

func (g *AuthenticationGateway) Generate(userID string) (string, error) {
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

func (g *AuthenticationGateway) GetUserID(tokenString string) (string, error) {
	claims, err := g.verify(tokenString)
	if err != nil {
		return "", err
	}

	return claims.UserID, err
}

func (g *AuthenticationGateway) GetUserIDBypassTokenExpiry(tokenString string) (string, error) {
	claims, err := g.verifyBypassExpiry(tokenString)
	if err != nil {
		return "", err
	}

	return claims.UserID, err
}

func (g *AuthenticationGateway) GetIssuedAt(tokenString string) (*time.Time, error) {
	claims, err := g.verifyBypassExpiry(tokenString)
	if err != nil {
		return nil, err
	}

	issuedAt := claims.IssuedAt.Time

	return &issuedAt, nil
}

func NewAuthenticationGateway(
	config *configs.Config,
) (*AuthenticationGateway, error) {
	return &AuthenticationGateway{
		config: AuthenticationConfig{
			secret: []byte(config.JWT.Secret),
		},
	}, nil
}
