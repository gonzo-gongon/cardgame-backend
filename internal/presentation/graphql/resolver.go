package graphql

import (
	"original-card-game-backend/internal/application/usecase"
	"original-card-game-backend/internal/presentation/graphql/loader"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	loaders     *loader.Loaders
	cardUsecase *usecase.CardUsecase
}

func NewResolver(
	loaders *loader.Loaders,
	cardUsecase *usecase.CardUsecase,
) *Resolver {
	return &Resolver{
		loaders:     loaders,
		cardUsecase: cardUsecase,
	}
}
