package model

import (
	"original-card-game-backend/internal/infrastructure/value"
	"time"
)

type UserSession struct {
	General         `gorm:"embedded"`
	UserID          value.UUID[User] `gorm:"primarykey;column:user_id;type:binary(16);"`
	LatestSessionAt time.Time        `gorm:"column:latest_session_at;type:datetime;"`
}
