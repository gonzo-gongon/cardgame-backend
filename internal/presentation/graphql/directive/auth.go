package directive

import (
	"context"
	"fmt"
	"original-card-game-backend/internal/application/usecase"
	"original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/presentation/middleware"

	"github.com/99designs/gqlgen/graphql"
)

type AuthDirective struct {
	authenticationUsecase *usecase.AuthenticationUsecase
}

func NewAuthDirective(
	authenticationUsecase *usecase.AuthenticationUsecase,
) (*AuthDirective, error) {
	return &AuthDirective{
		authenticationUsecase: authenticationUsecase,
	}, nil
}

const userContextKey = "user"

func (d *AuthDirective) Auth(
	ctx context.Context,
	obj interface{},
	next graphql.Resolver,
) (interface{}, error) {
	gctx, err := middleware.ContextToGinContext(ctx)

	if err != nil {
		return false, err
	}

	token := middleware.GetToken(gctx)

	if token == "" {
		return false, fmt.Errorf("no token supplied")
	}

	user, err := d.authenticationUsecase.GetUser(token)
	if err != nil {
		return false, err
	}

	ctx = context.WithValue(ctx, userContextKey, user)

	return next(ctx)
}

func GetUserFromContext(ctx context.Context) (*model.User, error) {
	c := ctx.Value(userContextKey)
	user := c.(*model.User)
	if user == nil {
		return nil, fmt.Errorf("no user specified")
	}

	return user, nil
}
