package loader

import (
	"original-card-game-backend/internal/presentation/graphql/model"

	"github.com/graph-gophers/dataloader/v7"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

func NewLoaders(
	userLoader *UserLoader,
) *Loaders {
	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userLoader.BatchGetUsers),
	}
}
