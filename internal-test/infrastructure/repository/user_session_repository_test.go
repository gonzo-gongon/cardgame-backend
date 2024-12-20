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
	time "time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"go.uber.org/mock/gomock"
)

func TestUserSessionRepository_GetUpdatedAt_正常系(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	domainUserID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	var userID value.UUID[inframodel.User]
	assert.NoError(t, (&userID).Parse(domainUserID.String()))

	loc, _ := time.LoadLocation("Asia/Tokyo")
	latestSessionAt := time.Date(2024, 12, 14, 0, 0, 0, 0, loc)

	conn, mock, _ := mockgateway.NewDBMock()

	rows := sqlmock.NewRows([]string{"user_id", "latest_session_at"}).
		AddRow(userID, latestSessionAt)

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `user_sessions` WHERE `user_sessions`.`user_id` = ?")).
		WithArgs(userID).
		WillReturnRows(rows)

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
	assert.NoError(t, container.Provide(repository.NewUserSessionRepository))

	assert.NoError(t, container.Invoke(func(r repository.UserSessionRepository) {
		actual, actualError := r.GetUpdatedAt(domainUserID.String())
		assert.NoError(t, actualError)
		assert.Equal(t, &latestSessionAt, actual)
	}))
}
func TestUserSessionRepository_Create_正常系(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	domainUserID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	var userID value.UUID[inframodel.User]
	assert.NoError(t, (&userID).Parse(domainUserID.String()))

	loc, _ := time.LoadLocation("Asia/Tokyo")
	latestSessionAt := time.Date(2024, 12, 14, 0, 0, 0, 0, loc)

	conn, mock, _ := mockgateway.NewDBMock()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO `user_sessions` (`created_at`,`updated_at`,`user_id`,`latest_session_at`) VALUES (?,?,?,?)")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), userID, latestSessionAt).
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
	assert.NoError(t, container.Provide(repository.NewUserSessionRepository))

	assert.NoError(t, container.Invoke(func(r repository.UserSessionRepository) {
		actualError := r.Create(domainUserID.String(), &latestSessionAt)
		assert.NoError(t, actualError)
	}))
}

func TestUserSessionRepository_Update_正常系(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	domainUserID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	var userID value.UUID[inframodel.User]
	assert.NoError(t, (&userID).Parse(domainUserID.String()))

	loc, _ := time.LoadLocation("Asia/Tokyo")
	latestSessionAt := time.Date(2024, 12, 14, 0, 0, 0, 0, loc)

	conn, mock, _ := mockgateway.NewDBMock()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(
		"UPDATE `user_sessions` SET `updated_at`=?,`latest_session_at`=? WHERE `user_id` = ?")).
		WithArgs(sqlmock.AnyArg(), latestSessionAt, userID).
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
	assert.NoError(t, container.Provide(repository.NewUserSessionRepository))

	assert.NoError(t, container.Invoke(func(r repository.UserSessionRepository) {
		actualError := r.Update(domainUserID.String(), &latestSessionAt)
		assert.NoError(t, actualError)
	}))
}
