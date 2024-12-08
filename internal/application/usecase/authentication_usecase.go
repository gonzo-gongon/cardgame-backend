package usecase

import (
	"fmt"
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/repository"
)

type AuthenticationUsecase struct {
	authenticationRepository *repository.AuthenticationRepository
	userRepository           repository.IUserRepository
	userSessionRepository    *repository.UserSessionRepository
}

func (u *AuthenticationUsecase) SignUp() (string, error) {
	createUser := repository.CreateUser{
		Name: "name",
	}
	user, err := u.userRepository.Create(&createUser)
	if err != nil {
		return "", err
	}

	token, err := u.authenticationRepository.Generate(user.ID)
	if err != nil {
		return "", err
	}

	issuedAt, err := u.authenticationRepository.GetIssuedAt(token)
	if err != nil {
		return "", err
	}

	err = u.userSessionRepository.Create(user.ID.String(), issuedAt)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *AuthenticationUsecase) Refresh(token string) (string, error) {
	issuedAt, err := u.authenticationRepository.GetIssuedAt(token)
	if err != nil {
		return "", err
	}

	userID, err := u.authenticationRepository.GetUserIDBypassTokenExpiry(token)
	if err != nil {
		return "", err
	}

	userIDValue := *userID

	updatedAt, err := u.userSessionRepository.GetUpdatedAt(userIDValue.String())
	if err != nil {
		return "", err
	}

	if updatedAt.Compare(*issuedAt) != 0 {
		fmt.Printf("t1: %s t2: %s\n", issuedAt, updatedAt)
		return "", fmt.Errorf("this is not the latest token")
	}

	newToken, err := u.authenticationRepository.Generate(userIDValue)
	if err != nil {
		return "", err
	}

	newIssuedAt, err := u.authenticationRepository.GetIssuedAt(newToken)
	if err != nil {
		return "", err
	}

	err = u.userSessionRepository.Update(userIDValue.String(), newIssuedAt)
	if err != nil {
		return "", err
	}

	return newToken, nil
}

func (u *AuthenticationUsecase) GetUser(token string) (*model.User, error) {
	userID, err := u.authenticationRepository.GetUserID(token)
	if err != nil {
		return nil, err
	}

	user, err := u.userRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewAuthenticationUsecase(
	authenticationRepository *repository.AuthenticationRepository,
	userRepository repository.IUserRepository,
	userSessionRepository *repository.UserSessionRepository,
) (*AuthenticationUsecase, error) {
	return &AuthenticationUsecase{
		authenticationRepository: authenticationRepository,
		userRepository:           userRepository,
		userSessionRepository:    userSessionRepository,
	}, nil
}
