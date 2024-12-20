package model

import (
	"time"

	"gorm.io/gorm"
)

type General struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type GeneralWithDelete struct {
	General   `gorm:"embedded" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
