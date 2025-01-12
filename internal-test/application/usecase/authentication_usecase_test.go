package usecase_test

import (
	mockrepository "original-card-game-backend/internal-test/infrastructure/repository"
	"original-card-game-backend/internal/application/usecase"
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"go.uber.org/mock/gomock"
)

type NotLatestTokenError struct{}

func (e *NotLatestTokenError) Error() string {
	return "this is not the latest token"
}

type TokenNotFoundError struct{}

func (e *TokenNotFoundError) Error() string {
	return "token not found"
}

func TestAuthenticationUsecase_SignUp_正常系(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	const token = "thisistoken"

	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	user := &model.User{
		ID:   userID,
		Name: "name",
	}
	createUser := repository.CreateUser{
		Name: "name",
	}
	loc, _ := time.LoadLocation("Asia/Tokyo")
	issuedAt := time.Date(2024, 12, 14, 0, 0, 0, 0, loc)

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.AuthenticationRepository {
		r := mockrepository.NewMockAuthenticationRepository(c)
		r.EXPECT().Generate(userID).Return(token, nil)
		r.EXPECT().GetIssuedAt(token).Return(&issuedAt, nil)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserRepository {
		r := mockrepository.NewMockUserRepository(c)
		r.EXPECT().Create(&createUser).Return(user, nil)
		assert.Equal(t, user.Name, createUser.Name)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserSessionRepository {
		r := mockrepository.NewMockUserSessionRepository(c)
		r.EXPECT().Create(userID.String(), &issuedAt).Return(nil)

		return r
	}))
	assert.NoError(t, container.Provide(usecase.NewAuthenticationUsecase))

	assert.NoError(t, container.Invoke(func(u *usecase.AuthenticationUsecase) {
		actual, actualError := u.SignUp()
		assert.NoError(t, actualError)
		assert.Equal(t, token, actual)
	}))
}

func TestAuthenticationUsecase_Refresh_正常系(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	const token = "thisistoken"

	const newToken = "thisisnewtoken"

	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	loc, _ := time.LoadLocation("Asia/Tokyo")
	issuedAt := time.Date(2024, 12, 14, 0, 0, 0, 0, loc)
	newIssuedAt := time.Date(2024, 12, 14, 1, 0, 0, 0, loc)
	updatedAt := time.Date(2024, 12, 14, 0, 0, 0, 0, loc)

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	//nolint:varnamelen // テストコードのため許容する
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.AuthenticationRepository {
		r := mockrepository.NewMockAuthenticationRepository(c)
		r.EXPECT().Generate(userID).Return(newToken, nil)
		r.EXPECT().GetIssuedAt(token).Return(&issuedAt, nil)
		r.EXPECT().GetIssuedAt(newToken).Return(&newIssuedAt, nil)
		r.EXPECT().GetUserIDBypassTokenExpiry(token).Return(&userID, nil)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserRepository {
		r := mockrepository.NewMockUserRepository(c)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserSessionRepository {
		r := mockrepository.NewMockUserSessionRepository(c)
		r.EXPECT().GetUpdatedAt(userID.String()).Return(&updatedAt, nil)
		r.EXPECT().Update(userID.String(), &newIssuedAt).Return(nil)

		return r
	}))
	assert.NoError(t, container.Provide(usecase.NewAuthenticationUsecase))

	assert.NoError(t, container.Invoke(func(u *usecase.AuthenticationUsecase) {
		actual, actualError := u.Refresh(token)
		assert.NoError(t, actualError)
		assert.Equal(t, newToken, actual)
	}))
}

func TestAuthenticationUsecase_Refresh_異常系_トークンが最新ではない(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	const token = "thisistoken"

	const newToken = "thisisnewtoken"

	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	loc, _ := time.LoadLocation("Asia/Tokyo")
	issuedAt := time.Date(2024, 12, 13, 0, 0, 0, 0, loc)
	updatedAt := time.Date(2024, 12, 14, 0, 0, 0, 0, loc)

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	//nolint:varnamelen // this is testcode
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.AuthenticationRepository {
		r := mockrepository.NewMockAuthenticationRepository(c)
		r.EXPECT().Generate(gomock.Any()).Times(0)
		r.EXPECT().GetIssuedAt(token).Return(&issuedAt, nil)
		r.EXPECT().GetIssuedAt(newToken).Times(0)
		r.EXPECT().GetUserIDBypassTokenExpiry(token).Return(&userID, nil)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserRepository {
		r := mockrepository.NewMockUserRepository(c)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserSessionRepository {
		r := mockrepository.NewMockUserSessionRepository(c)
		r.EXPECT().GetUpdatedAt(userID.String()).Return(&updatedAt, nil)
		r.EXPECT().Update(gomock.Any(), gomock.Any()).Times(0)

		return r
	}))
	assert.NoError(t, container.Provide(usecase.NewAuthenticationUsecase))

	assert.NoError(t, container.Invoke(func(u *usecase.AuthenticationUsecase) {
		actual, actualError := u.Refresh(token)
		assert.Error(t, actualError, NotLatestTokenError{})
		assert.Equal(t, "", actual)
	}))
}

func TestAuthenticationUsecase_GetUser_正常系(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	const token = "thisistoken"

	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	expected := &model.User{
		ID:   userID,
		Name: "",
	}

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.AuthenticationRepository {
		r := mockrepository.NewMockAuthenticationRepository(c)
		r.EXPECT().GetUserID(gomock.Any()).Return(userID.String(), nil)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserRepository {
		r := mockrepository.NewMockUserRepository(c)
		r.EXPECT().GetByUserID(gomock.Any()).Return(expected, nil)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserSessionRepository {
		r := mockrepository.NewMockUserSessionRepository(c)

		return r
	}))
	assert.NoError(t, container.Provide(usecase.NewAuthenticationUsecase))

	assert.NoError(t, container.Invoke(func(u *usecase.AuthenticationUsecase) {
		actual, actualError := u.GetUser(token)
		assert.NoError(t, actualError)
		assert.Equal(t, expected, actual)
	}))
}

func TestAuthenticationUsecase_GetUser_異常系_該当ユーザーなし(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	const token = "thisistoken"

	var expected *model.User

	expectedError := &TokenNotFoundError{}

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))

	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.AuthenticationRepository {
		r := mockrepository.NewMockAuthenticationRepository(c)
		r.EXPECT().GetUserID(gomock.Any()).Return("", expectedError)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserRepository {
		r := mockrepository.NewMockUserRepository(c)
		r.EXPECT().GetByUserID(gomock.Any()).Times(0)

		return r
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserSessionRepository {
		r := mockrepository.NewMockUserSessionRepository(c)

		return r
	}))
	assert.NoError(t, container.Provide(usecase.NewAuthenticationUsecase))

	assert.NoError(t, container.Invoke(func(u *usecase.AuthenticationUsecase) {
		actual, actualError := u.GetUser(token)
		assert.ErrorIs(t, actualError, expectedError)
		assert.Equal(t, expected, actual)
	}))
}
