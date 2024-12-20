package repository

import (
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/gateway"
	inframodel "original-card-game-backend/internal/infrastructure/model"
	"original-card-game-backend/internal/infrastructure/value"
)

type UserNotFoundError struct{}

func (e *UserNotFoundError) Error() string {
	return "user not found error"
}

type UserRepository interface {
	GetByUserID(userID string) (*model.User, error)
	Create(createUser *CreateUser) (*model.User, error)
}

type UserRepositoryImpl struct {
	databaseGateway gateway.DatabaseGateway
}

type CreateUser struct {
	Name string
}

func (r *UserRepositoryImpl) GetByUserID(userID string) (*model.User, error) {
	conn, err := r.databaseGateway.Connect()
	if err != nil {
		return nil, err
	}

	var parsedUserID value.UUID[inframodel.User]
	err = parsedUserID.Parse(userID)

	if err != nil {
		return nil, err
	}

	user := inframodel.User{
		ID: parsedUserID,
	}

	result := conn.Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, &UserNotFoundError{}
	}

	entity := user.Domain()

	return &entity, nil
}

func (r *UserRepositoryImpl) Create(createUser *CreateUser) (*model.User, error) {
	conn, err := r.databaseGateway.Connect()
	if err != nil {
		return nil, err
	}

	user := inframodel.User{
		Name: createUser.Name,
	}
	if result := conn.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	ret := user.Domain()

	return &ret, nil
}

//nolint:ireturn // DIのためのコードなので許容する
func NewUserRepository(
	databaseGateway gateway.DatabaseGateway,
) (UserRepository, error) {
	return &UserRepositoryImpl{
		databaseGateway: databaseGateway,
	}, nil
}
