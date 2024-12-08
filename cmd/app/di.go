package app

import (
	"original-card-game-backend/configs"
	"original-card-game-backend/internal/application/usecase"
	"original-card-game-backend/internal/infrastructure/gateway"
	"original-card-game-backend/internal/infrastructure/repository"
	"original-card-game-backend/internal/presentation/controller"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()

	// if err := container.Provide(func() *configs.Config {
	// 	return configs.NewConfigs()
	// }); err != nil {
	// 	panic(err)
	// }

	// if err := container.Provide(func(cfg *configs.Config) (*gateway.DatabaseGateway, error) {
	// 	return gateway.NewDatabaseGateway(cfg)
	// }); err != nil {
	// 	panic(err)
	// }

	// Config
	if err := container.Provide(configs.NewConfigs); err != nil {
		panic(err)
	}

	// Gateway
	if err := container.Provide(gateway.NewAuthenticationGateway); err != nil {
		panic(err)
	}

	if err := container.Provide(gateway.NewDatabaseGateway); err != nil {
		panic(err)
	}

	// Repository
	if err := container.Provide(repository.NewAuthenticationRepository); err != nil {
		panic(err)
	}

	if err := container.Provide(repository.NewUserRepository); err != nil {
		panic(err)
	}

	if err := container.Provide(repository.NewUserSessionRepository); err != nil {
		panic(err)
	}

	// Usecase
	if err := container.Provide(usecase.NewAuthenticationUsecase); err != nil {
		panic(err)
	}

	if err := container.Provide(usecase.NewUserUsecase); err != nil {
		panic(err)
	}

	// Controller
	if err := container.Provide(controller.NewAuthenticationController); err != nil {
		panic(err)
	}

	if err := container.Provide(controller.NewUserController); err != nil {
		panic(err)
	}

	return container
}
