package repository

import (
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/gateway"
	inframodel "original-card-game-backend/internal/infrastructure/model"
	"original-card-game-backend/internal/infrastructure/value"
)

type IUserRepository interface {
	GetByUserID(userID string) (*model.User, error)
	Create(createUser *CreateUser) (*model.User, error)
}

type UserRepository struct {
	databaseGateway *gateway.DatabaseGateway
}

type CreateUser struct {
	Name string
}

func (r *UserRepository) GetByUserID(userID string) (*model.User, error) {
	db, err := r.databaseGateway.Connect()
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

	if result := db.Find(&user); result.Error != nil {
		return nil, result.Error
	}

	entity := user.Domain()

	return &entity, nil
}

func (r *UserRepository) Create(createUser *CreateUser) (*model.User, error) {
	db, err := r.databaseGateway.Connect()
	if err != nil {
		return nil, err
	}

	user := inframodel.User{
		Name: createUser.Name,
	}
	if result := db.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	ret := user.Domain()

	return &ret, nil
}

func NewUserRepository(
	databaseGateway *gateway.DatabaseGateway,
) (IUserRepository, error) {
	return &UserRepository{
		databaseGateway: databaseGateway,
	}, nil
}
