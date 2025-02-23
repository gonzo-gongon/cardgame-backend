package app

import (
	"original-card-game-backend/configs"
	"original-card-game-backend/internal/application/usecase"
	"original-card-game-backend/internal/infrastructure/gateway"
	"original-card-game-backend/internal/infrastructure/repository"
	"original-card-game-backend/internal/presentation/controller"
	"original-card-game-backend/internal/presentation/graphql"
	"original-card-game-backend/internal/presentation/graphql/directive"
	"original-card-game-backend/internal/presentation/graphql/loader"
	"original-card-game-backend/internal/presentation/middleware"
	"original-card-game-backend/internal/presentation/presenter"

	"go.uber.org/dig"
)

func BuildContainer() *dig.Container { //nolint:cyclop // DIのためのコードなので許容する
	container := dig.New()

	// Config
	if err := container.Provide(configs.NewConfigs); err != nil {
		panic(err)
	}

	// GraphQL
	if err := container.Provide(loader.NewLoaders); err != nil {
		panic(err)
	}

	if err := container.Provide(loader.NewUserLoader); err != nil {
		panic(err)
	}

	if err := container.Provide(graphql.NewResolver); err != nil {
		panic(err)
	}

	if err := container.Provide(directive.NewAuthDirective); err != nil {
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

	if err := container.Provide(repository.NewCardRepository); err != nil {
		panic(err)
	}

	// Usecase
	if err := container.Provide(usecase.NewAuthenticationUsecase); err != nil {
		panic(err)
	}

	if err := container.Provide(usecase.NewUserUsecase); err != nil {
		panic(err)
	}

	if err := container.Provide(usecase.NewCardUsecase); err != nil {
		panic(err)
	}

	// Controller
	if err := container.Provide(controller.NewAuthenticationController); err != nil {
		panic(err)
	}
	if err := container.Provide(controller.NewGraphQLController); err != nil {
		panic(err)
	}

	// Middleware
	if err := container.Provide(middleware.NewTokenRefreshMiddleware); err != nil {
		panic(err)
	}

	if err := container.Provide(middleware.NewAuthenticationMiddleware); err != nil {
		panic(err)
	}

	if err := container.Provide(middleware.NewContextMiddleware); err != nil {
		panic(err)
	}

	// Presenter
	if err := container.Provide(presenter.NewAuthenticationPresenter); err != nil {
		panic(err)
	}

	return container
}
