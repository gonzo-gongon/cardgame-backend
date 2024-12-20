package value

import (
	"database/sql/driver"
	"original-card-game-backend/internal/domain/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type InvalidUUIDError struct{}

func (e *InvalidUUIDError) Error() string {
	return "invalid uuid"
}

type UUID[T any] uuid.UUID

func (u *UUID[T]) New() UUID[T] {
	uid := uuid.Must(uuid.NewV7())

	return UUID[T](uid)
}

func (u *UUID[T]) GormDataType() string {
	return "binary(16)"
}

func (u *UUID[T]) GormDBDataType(_ *gorm.DB, _ *schema.Field) string { //nolint:revive // 引数未使用だがgorm側で呼び出すためそのままにしておく
	return "binary"
}

func (u *UUID[T]) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return &InvalidUUIDError{}
	}

	parseByte, err := uuid.FromBytes(bytes)
	*u = UUID[T](parseByte)

	return err
}

func (u UUID[T]) Value() (driver.Value, error) {
	return uuid.UUID(u).MarshalBinary()
}

func (u UUID[T]) String() string {
	return uuid.UUID(u).String()
}

func (u *UUID[T]) Parse(str string) error {
	p, err := uuid.Parse(str)
	if err != nil {
		return &InvalidUUIDError{}
	}

	*u = UUID[T](p)

	return nil
}

func UUIDToDomain[T, U any](u UUID[T]) model.UUID[U] {
	return model.UUID[U](u.String())
}
