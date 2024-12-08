package usecase

import (
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/repository"
)

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func (u *UserUsecase) GetUser(userID string) (*model.User, error) {
	user, err := u.userRepository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserUsecase(
	userRepository repository.IUserRepository,
) (*UserUsecase, error) {
	return &UserUsecase{
		userRepository: userRepository,
	}, nil
}
