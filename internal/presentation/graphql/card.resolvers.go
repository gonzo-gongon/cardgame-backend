package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.63

import (
	"context"
	"errors"
	domain "original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/presentation/graphql/core"
	"original-card-game-backend/internal/presentation/graphql/directive"
	"original-card-game-backend/internal/presentation/graphql/model"
)

// CreatedBy is the resolver for the createdBy field.
func (r *cardResolver) CreatedBy(ctx context.Context, obj *model.Card) (*model.User, error) {
	if obj.CreatedBy == nil {
		return nil, errors.New("not found")
	}

	load := r.loaders.UserLoader.Load(ctx, obj.CreatedBy.ID)

	user, err := load()

	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdatedBy is the resolver for the updatedBy field.
func (r *cardResolver) UpdatedBy(ctx context.Context, obj *model.Card) (*model.User, error) {
	if obj.UpdatedBy == nil {
		return nil, errors.New("not found")
	}

	load := r.loaders.UserLoader.Load(ctx, obj.UpdatedBy.ID)

	user, err := load()

	if err != nil {
		return nil, err
	}

	return user, nil
}

// CreateCard is the resolver for the createCard field.
func (r *mutationResolver) CreateCard(ctx context.Context, input model.CreateCardInput) (*model.Card, error) {
	user, err := directive.GetUserFromContext(ctx)

	if err != nil {
		return nil, err
	}

	ret, err := r.cardUsecase.CreateCard(
		domain.CreateCard{
			Name: input.Name,
			Text: input.Text,
		},
		user,
	)

	if err != nil {
		panic(err)
	}

	return &model.Card{
		ID:   ret.ID.String(),
		Name: ret.Name,
		Text: ret.Text,
		CreatedBy: &model.User{
			ID: ret.CreatedBy.String(),
		},
		UpdatedBy: &model.User{
			ID: ret.UpdatedBy.String(),
		},
	}, nil
}

// Cards is the resolver for the cards field.
func (r *queryResolver) Cards(ctx context.Context, ids []string) ([]*model.Card, error) {
	cards, err := r.cardUsecase.GetCards(domain.UUIDsFromString[domain.Card](ids))

	if err != nil {
		panic(err)
	}

	ret := make([]*model.Card, len(cards))

	for i, v := range cards {
		ret[i] = &model.Card{
			ID:   string(v.ID),
			Name: v.Name,
			Text: v.Text,
			CreatedBy: &model.User{
				ID: v.CreatedBy.String(),
			},
			UpdatedBy: &model.User{
				ID: v.UpdatedBy.String(),
			},
		}
	}

	return ret, nil
}

// Card returns core.CardResolver implementation.
func (r *Resolver) Card() core.CardResolver { return &cardResolver{r} }

type cardResolver struct{ *Resolver }
