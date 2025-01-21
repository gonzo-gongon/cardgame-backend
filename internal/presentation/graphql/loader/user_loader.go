package loader

import (
	"context"
	"original-card-game-backend/internal/application/usecase"
	domain "original-card-game-backend/internal/domain/model"
	"original-card-game-backend/internal/presentation/graphql/model"

	"github.com/graph-gophers/dataloader/v7"
)

type UserLoader struct {
	userUsecase *usecase.UserUsecase
}

func (l *UserLoader) BatchGetUsers(ctx context.Context, IDs []string) []*dataloader.Result[*model.User] {
	userIDs := domain.UUIDsFromString[domain.User](IDs)

	users, err := l.userUsecase.GetUsersByIDs(userIDs)

	indexMapper := make(map[domain.UUID[domain.User]]int, len(userIDs))

	for i, ID := range userIDs {
		indexMapper[ID] = i
	}

	results := make([]*dataloader.Result[*model.User], len(IDs))

	for _, user := range users {
		if err != nil {
			results[indexMapper[user.ID]] = &dataloader.Result[*model.User]{
				Error: err,
			}
		}

		results[indexMapper[user.ID]] = &dataloader.Result[*model.User]{
			Data: &model.User{
				ID:   user.ID.String(),
				Name: user.Name,
			},
		}
	}

	return results
}

func NewUserLoader(
	userUsecase *usecase.UserUsecase,
) (*UserLoader, error) {
	return &UserLoader{
		userUsecase: userUsecase,
	}, nil
}
