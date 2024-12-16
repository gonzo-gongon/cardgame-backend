package repository

import (
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/gateway"
	"time"
)

type AuthenticationRepository interface {
	Generate(userID model.UUID[model.User]) (string, error)
	GetUserIDBypassTokenExpiry(tokenString string) (*model.UUID[model.User], error)
	GetUserID(tokenString string) (string, error)
	GetIssuedAt(tokenString string) (*time.Time, error)
}

type AuthenticationRepositoryImpl struct {
	authenticationGateway gateway.AuthenticationGateway
}

func (r *AuthenticationRepositoryImpl) Generate(userID model.UUID[model.User]) (string, error) {
	return r.authenticationGateway.Generate(userID.String())
}

func (r *AuthenticationRepositoryImpl) GetUserIDBypassTokenExpiry(tokenString string) (*model.UUID[model.User], error) {
	userID, err := r.authenticationGateway.GetUserIDBypassTokenExpiry(tokenString)

	if err != nil {
		return nil, err
	}

	parsedUserID := model.UUID[model.User](userID)

	return &parsedUserID, err
}

func (r *AuthenticationRepositoryImpl) GetUserID(tokenString string) (string, error) {
	return r.authenticationGateway.GetUserID(tokenString)
}

func (r *AuthenticationRepositoryImpl) GetIssuedAt(tokenString string) (*time.Time, error) {
	return r.authenticationGateway.GetIssuedAt(tokenString)
}

//nolint:ireturn
func NewAuthenticationRepository(
	authenticationGateway gateway.AuthenticationGateway,
) (AuthenticationRepository, error) {
	return &AuthenticationRepositoryImpl{
		authenticationGateway: authenticationGateway,
	}, nil
}
