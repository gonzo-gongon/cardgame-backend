package gateway

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MockCreationError struct {
	cause error
}

func (e *MockCreationError) Error() string {
	return fmt.Sprintf("mock creation error: %s", e.cause.Error())
}

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
