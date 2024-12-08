package repository

import (
	"original-card-game-backend/internal/infrastructure/gateway"
	"original-card-game-backend/internal/infrastructure/model"
	"original-card-game-backend/internal/infrastructure/value"
	"time"
)

type UserSessionRepository struct {
	databaseGateway *gateway.DatabaseGateway
}

func (r *UserSessionRepository) GetUpdatedAt(userID string) (*time.Time, error) {
	db, err := r.databaseGateway.Connect()
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
	if result := db.Find(&userSession); result.Error != nil {
		return nil, result.Error
	}

	return &userSession.LatestSessionAt, nil
}

func (r *UserSessionRepository) Create(userID string, createdAt *time.Time) error {
	db, err := r.databaseGateway.Connect()
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
	if result := db.Create(&userSession); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserSessionRepository) Update(userID string, updatedAt *time.Time) error {
	db, err := r.databaseGateway.Connect()
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
	if result := db.Model(&userSession).Updates(model.UserSession{
		LatestSessionAt: *updatedAt,
	}); result.Error != nil {
		return result.Error
	}

	return nil
}

func NewUserSessionRepository(
	databaseGateway *gateway.DatabaseGateway,
) (*UserSessionRepository, error) {
	return &UserSessionRepository{
		databaseGateway: databaseGateway,
	}, nil
}
