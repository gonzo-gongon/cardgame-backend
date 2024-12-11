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

func TestAuthenticationUsecase_GetUser_正常系(t *testing.T) {
	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")
	token := "thisistoken"
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

func TestAuthenticationUsecase_GetUser_異常系_該当ユーザーなし(t *testing.T) {
	token := "thisistoken"
	var expected *model.User
	expectedError := errors.New("token not found")

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
		assert.Error(t, actualError, expectedError)
		assert.Equal(t, expected, actual)
	}))
}
