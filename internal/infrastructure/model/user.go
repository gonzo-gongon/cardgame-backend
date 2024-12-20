package model

import (
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/value"

	"gorm.io/gorm"
)

type User struct {
	GeneralWithDelete `gorm:"embedded"`
	ID                value.UUID[User] `gorm:"primarykey;column:id;type:binary(16);<-:create"`
	Name              string           `gorm:"column:name;type:varchar(30);"`
}

func (d *User) BeforeCreate(db *gorm.DB) (err error) { //nolint:revive,nonamedreturns // 引数未使用だがgorm側で呼び出すためそのままにしておく
	d.ID = d.ID.New()

	return
}

func (d *User) Domain() model.User {
	return model.User{
		ID:   value.UUIDToDomain[User, model.User](d.ID),
		Name: d.Name,
	}
}
