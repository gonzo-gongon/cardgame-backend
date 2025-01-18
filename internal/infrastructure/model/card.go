package model

import (
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/value"

	"gorm.io/gorm"
)

type Card struct {
	GeneralWithDelete `gorm:"embedded"`
	ID                value.UUID[Card]  `gorm:"primarykey;column:id;type:binary(16);<-:create"`
	Name              string            `gorm:"column:name;type:varchar(30);"`
	Text              string            `gorm:"column:text;type:mediumtext;"`
	CreatedBy         *value.UUID[User] `gorm:"column:created_by;type:binary(16);"`
	UpdatedBy         *value.UUID[User] `gorm:"column:updated_by;type:binary(16);"`
	DeletedBy         *value.UUID[User] `gorm:"column:deleted_by;type:binary(16);"`
}

type Cards []Card

func (d *Card) BeforeCreate(db *gorm.DB) (err error) { //nolint:revive,nonamedreturns // 引数未使用だがgorm側で呼び出すためそのままにしておく
	d.ID = d.ID.New()

	return
}

func (d *Card) Domain() model.Card {
	return model.Card{
		ID:   value.UUIDToDomain[Card, model.Card](d.ID),
		Name: d.Name,
		Text: d.Text,
	}
}

func (d Cards) Domain() []model.Card {
	r := make([]model.Card, len(d))

	for i, v := range d {
		r[i] = v.Domain()
	}

	return r
}
