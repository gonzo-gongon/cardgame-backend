package graphql

import "original-card-game-backend/internal/application/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	cardUsecase *usecase.CardUsecase
}

func NewResolver(
	cardUsecase *usecase.CardUsecase,
) *Resolver {
	return &Resolver{
		cardUsecase: cardUsecase,
	}
}
