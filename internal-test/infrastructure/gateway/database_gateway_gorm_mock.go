package gateway

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//nolint:ireturn // DIのためのコードなので許容する
func NewDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()

	if err != nil {
		return nil, nil, err
	}

	//nolint:exhaustruct // テストのためのモックコードなので許容する
	mockDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	return mockDB, mock, err
}
