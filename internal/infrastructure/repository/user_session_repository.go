package repository

import (
	"original-card-game-backend/internal/infrastructure/gateway"
	"original-card-game-backend/internal/infrastructure/model"
	"original-card-game-backend/internal/infrastructure/value"
	"time"
)

type UserSessionRepository interface {
	GetUpdatedAt(userID string) (*time.Time, error)
	Create(userID string, createdAt *time.Time) error
	Update(userID string, updatedAt *time.Time) error
}

type UserSessionRepositoryImpl struct {
	databaseGateway gateway.DatabaseGateway
}

func (r *UserSessionRepositoryImpl) GetUpdatedAt(userID string) (*time.Time, error) {
	conn, err := r.databaseGateway.Connect()
	if err != nil {
		return nil, err
	}

	var parsedUserID value.UUID[model.User]
	err = parsedUserID.Parse(userID)

	if err != nil {
		return nil, err
	}

	userSession := model.UserSession{
		UserID: parsedUserID,
	}
	if result := conn.Find(&userSession); result.Error != nil {
		return nil, result.Error
	}

	return &userSession.LatestSessionAt, nil
}

func (r *UserSessionRepositoryImpl) Create(userID string, createdAt *time.Time) error {
	conn, err := r.databaseGateway.Connect()
	if err != nil {
		return err
	}

	var parsedUserID value.UUID[model.User]
	err = parsedUserID.Parse(userID)

	if err != nil {
		return err
	}

	userSession := model.UserSession{
		UserID:          parsedUserID,
		LatestSessionAt: *createdAt,
	}
	if result := conn.Create(&userSession); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserSessionRepositoryImpl) Update(userID string, updatedAt *time.Time) error {
	conn, err := r.databaseGateway.Connect()
	if err != nil {
		return err
	}

	var parsedUserID value.UUID[model.User]
	err = parsedUserID.Parse(userID)

	if err != nil {
		return err
	}

	userSession := model.UserSession{
		UserID: parsedUserID,
	}
	if result := conn.Model(&userSession).Updates(model.UserSession{
		LatestSessionAt: *updatedAt,
	}); result.Error != nil {
		return result.Error
	}

	return nil
}

//nolint:ireturn // DIのためのコードなので許容する
func NewUserSessionRepository(
	databaseGateway gateway.DatabaseGateway,
) (UserSessionRepository, error) {
	return &UserSessionRepositoryImpl{
		databaseGateway: databaseGateway,
	}, nil
}
