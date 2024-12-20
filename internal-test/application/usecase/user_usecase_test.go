package usecase_test

import (
	"errors"
	mockrepository "original-card-game-backend/internal-test/infrastructure/repository"
	"original-card-game-backend/internal/application/usecase"
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"go.uber.org/mock/gomock"
)

func TestUserUsecase_GetByUserID_正常系(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

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
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserRepository {
		r := mockrepository.NewMockUserRepository(c)
		r.EXPECT().GetByUserID(userID.String()).Return(expected, nil)

		return r
	}))
	assert.NoError(t, container.Provide(usecase.NewUserUsecase))

	assert.NoError(t, container.Invoke(func(u *usecase.UserUsecase) {
		actual, actualError := u.GetUser(userID.String())
		assert.NoError(t, actualError)
		assert.Equal(t, expected, actual)
	}))
}

func TestUserUsecase_GetByUserID_正常系_対象なし(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	var expected *model.User

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserRepository {
		r := mockrepository.NewMockUserRepository(c)
		r.EXPECT().GetByUserID(userID.String()).Return(expected, nil)

		return r
	}))
	assert.NoError(t, container.Provide(usecase.NewUserUsecase))

	assert.NoError(t, container.Invoke(func(u *usecase.UserUsecase) {
		actual, actualError := u.GetUser(userID.String())
		assert.NoError(t, actualError)
		assert.Equal(t, expected, actual)
	}))
}

func TestUserUsecase_GetByUserID_異常系_エラーあり(t *testing.T) { //nolint:asciicheck // テストメソッドのため許容する
	t.Parallel()

	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	var expected *model.User
	expectedError := errors.New("")

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) repository.UserRepository {
		r := mockrepository.NewMockUserRepository(c)
		r.EXPECT().GetByUserID(userID.String()).Return(expected, expectedError)

		return r
	}))
	assert.NoError(t, container.Provide(usecase.NewUserUsecase))

	assert.NoError(t, container.Invoke(func(u *usecase.UserUsecase) {
		actual, actualError := u.GetUser(userID.String())
		assert.ErrorIs(t, actualError, expectedError)
		assert.Equal(t, expected, actual)
	}))
}
