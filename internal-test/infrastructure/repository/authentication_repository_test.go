package repository_test

import (
	mockgateway "original-card-game-backend/internal-test/infrastructure/gateway"
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/infrastructure/gateway"
	"original-card-game-backend/internal/infrastructure/repository"
	"testing"
	time "time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
	"go.uber.org/mock/gomock"
)

func TestAuthenticationRepository_Generate_正常系(t *testing.T) {
	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")

	token := "thisistoken"

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) gateway.AuthenticationGateway {
		r := mockgateway.NewMockAuthenticationGateway(c)
		r.EXPECT().Generate(userID.String()).Return(token, nil)

		return r
	}))
	assert.NoError(t, container.Provide(repository.NewAuthenticationRepository))

	assert.NoError(t, container.Invoke(func(r repository.AuthenticationRepository) {
		actual, actualError := r.Generate(userID)
		assert.NoError(t, actualError)
		assert.Equal(t, token, actual)
	}))
}

func TestAuthenticationRepository_GetUserIDBypassTokenExpiry_正常系(t *testing.T) {
	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")

	token := "thisistoken"

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) gateway.AuthenticationGateway {
		r := mockgateway.NewMockAuthenticationGateway(c)
		r.EXPECT().GetUserIDBypassTokenExpiry(token).Return(userID.String(), nil)

		return r
	}))
	assert.NoError(t, container.Provide(repository.NewAuthenticationRepository))

	assert.NoError(t, container.Invoke(func(r repository.AuthenticationRepository) {
		actual, actualError := r.GetUserIDBypassTokenExpiry(token)
		assert.NoError(t, actualError)
		assert.Equal(t, &userID, actual)
	}))
}

func TestAuthenticationRepository_GetUserID_正常系(t *testing.T) {
	userID := model.UUID[model.User]("0193a685-4c73-7119-b5fb-ee3eb12f115a")

	token := "thisistoken"

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) gateway.AuthenticationGateway {
		r := mockgateway.NewMockAuthenticationGateway(c)
		r.EXPECT().GetUserID(token).Return(userID.String(), nil)

		return r
	}))
	assert.NoError(t, container.Provide(repository.NewAuthenticationRepository))

	assert.NoError(t, container.Invoke(func(r repository.AuthenticationRepository) {
		actual, actualError := r.GetUserID(token)
		assert.NoError(t, actualError)
		assert.Equal(t, userID.String(), actual)
	}))
}

func TestAuthenticationRepository_GetIssuedAt_正常系(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	issuedAt := time.Date(2024, 12, 14, 0, 0, 0, 0, loc)

	token := "thisistoken"

	container := dig.New()
	assert.NoError(t, container.Provide(func() *gomock.Controller {
		c := gomock.NewController(t)

		return c
	}))
	assert.NoError(t, container.Provide(func(c *gomock.Controller) gateway.AuthenticationGateway {
		r := mockgateway.NewMockAuthenticationGateway(c)
		r.EXPECT().GetIssuedAt(token).Return(&issuedAt, nil)

		return r
	}))
	assert.NoError(t, container.Provide(repository.NewAuthenticationRepository))

	assert.NoError(t, container.Invoke(func(r repository.AuthenticationRepository) {
		actual, actualError := r.GetIssuedAt(token)
		assert.NoError(t, actualError)
		assert.Equal(t, &issuedAt, actual)
	}))
}
