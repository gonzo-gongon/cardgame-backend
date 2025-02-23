package usecase

import (
	"fmt"
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/repository"
)

type UserUsecaseGetUserError struct {
	cause error
}

func (e *UserUsecaseGetUserError) Error() string {
	return fmt.Sprintf("user usecase get user failed: %o", e.cause)
}

type UserUsecase struct {
	userRepository repository.UserRepository
}

func (u *UserUsecase) GetUser(userID string) (*model.User, error) {
	user, err := u.userRepository.GetByUserID(userID)
	if err != nil {
		return nil, &UserUsecaseGetUserError{
			cause: err,
		}
	}

	return user, nil
}

func (u *UserUsecase) GetUsersByIDs(userIDs []model.UUID[model.User]) ([]model.User, error) {
	users, err := u.userRepository.GetUsersByUserIDs(userIDs)
	if err != nil {
		return nil, &UserUsecaseGetUserError{
			cause: err,
		}
	}

	return users, nil
}

func NewUserUsecase(
	userRepository repository.UserRepository,
) (*UserUsecase, error) {
	return &UserUsecase{
		userRepository: userRepository,
	}, nil
}
