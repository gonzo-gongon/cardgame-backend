package repository

import (
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/gateway"
	"time"
)

type AuthenticationRepository struct {
	authenticationGateway *gateway.AuthenticationGateway
}

func (r *AuthenticationRepository) Generate(userID model.UUID[model.User]) (string, error) {
	return r.authenticationGateway.Generate(userID.String())
}

func (r *AuthenticationRepository) GetUserIDBypassTokenExpiry(tokenString string) (*model.UUID[model.User], error) {
	userID, err := r.authenticationGateway.GetUserIDBypassTokenExpiry(tokenString)

	if err != nil {
		return nil, err
	}

	parsedUserId := model.UUID[model.User](userID)

	return &parsedUserId, err
}

func (r *AuthenticationRepository) GetUserID(tokenString string) (string, error) {
	return r.authenticationGateway.GetUserID(tokenString)
}

func (r *AuthenticationRepository) GetIssuedAt(tokenString string) (*time.Time, error) {
	return r.authenticationGateway.GetIssuedAt(tokenString)
}

func NewAuthenticationRepository(
	authenticationGateway *gateway.AuthenticationGateway,
) (*AuthenticationRepository, error) {
	return &AuthenticationRepository{
		authenticationGateway: authenticationGateway,
	}, nil
}
