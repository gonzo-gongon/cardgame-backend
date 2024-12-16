package repository_test

import (
	mockgateway "original-card-game-backend/internal-test/infrastructure/gateway"
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/gateway"
	inframodel "original-card-game-backend/internal/infrastructure/model"
	"original-card-game-backend/internal/infrastructure/repository"
	"original-card-game-backend/internal/infrastructure/value"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"go.uber.org/mock/gomock"
)

func TestUserRepository_GetByUserID_正常系_レコードあり(t *testing.T) { //nolint:asciicheck
	domainUserID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	var userID value.UUID[inframodel.User]
	assert.NoError(t, (&userID).Parse(domainUserID.String()))

	expected := &model.User{
		ID:   domainUserID,
		Name: "",
	}
	db, mock, _ := mockgateway.NewDBMock()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(userID, "")

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = ?")).
		WithArgs(userID).
		WillReturnRows(rows)

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) gateway.DatabaseGateway {
		r := mockgateway.NewMockDatabaseGateway(c)
		r.EXPECT().Connect().Return(db, nil)

		return r
	}))
	assert.NoError(t, container.Provide(repository.NewUserRepository))

	assert.NoError(t, container.Invoke(func(r repository.UserRepository) {
		actual, actualError := r.GetByUserID(domainUserID.String())
		assert.NoError(t, actualError)
		assert.Equal(t, expected, actual)
	}))
}

func TestUserRepository_GetByUserID_異常系_レコードなし(t *testing.T) { //nolint:asciicheck
	domainUserID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	var userID value.UUID[inframodel.User]
	assert.NoError(t, (&userID).Parse(domainUserID.String()))

	var expected *model.User
	db, mock, _ := mockgateway.NewDBMock()

	rows := sqlmock.NewRows([]string{"id", "name"})

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = ?")).
		WithArgs(userID).
		WillReturnRows(rows)

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) gateway.DatabaseGateway {
		r := mockgateway.NewMockDatabaseGateway(c)
		r.EXPECT().Connect().Return(db, nil)

		return r
	}))
	assert.NoError(t, container.Provide(repository.NewUserRepository))

	assert.NoError(t, container.Invoke(func(r repository.UserRepository) {
		actual, actualError := r.GetByUserID(domainUserID.String())
		assert.Error(t, actualError, &repository.UserNotFoundError{})
		assert.Equal(t, expected, actual)
	}))
}

// User.BeforeCreateでIDが強制挿入されるため、IDまわりの確認はしない
// User.BeforeCreateをスタブ化すれば回避できるが、そこまでメリットがなさそうなのでやらない
func TestUserRepository_Create_正常系(t *testing.T) { //nolint:asciicheck
	createUser := &repository.CreateUser{
		Name: "name",
	}

	expected := &model.User{
		Name: "name",
	}
	conn, mock, _ := mockgateway.NewDBMock()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`id`,`name`) VALUES (?,?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, sqlmock.AnyArg(), createUser.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) gateway.DatabaseGateway {
		r := mockgateway.NewMockDatabaseGateway(c)
		r.EXPECT().Connect().Return(conn, nil)

		return r
	}))
	assert.NoError(t, container.Provide(repository.NewUserRepository))

	assert.NoError(t, container.Invoke(func(r repository.UserRepository) {
		actual, actualError := r.Create(createUser)
		assert.NoError(t, actualError)

		// IDは比較対象から除外する
		expected.ID = actual.ID
		assert.Equal(t, expected, actual)
	}))
}
