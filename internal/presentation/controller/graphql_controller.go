package controller

import (
	graphql "original-card-game-backend/internal/presentation/graphql/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type GraphQLController struct {
	resolver graphql.ResolverRoot
}

func (c *GraphQLController) GraphQL(context *gin.Context) {
	h := handler.New(
		graphql.NewExecutableSchema(
			graphql.Config{
				Resolvers: c.resolver,
			},
		),
	)
	h.AddTransport(transport.POST{})

	//TODO: いったんここで定義する、必ず環境変数にする
	isDevelopmentEnvironment := true

	if isDevelopmentEnvironment {
		h.Use(extension.Introspection{})
	}

	h.ServeHTTP(context.Writer, context.Request)
}

func (c *GraphQLController) GraphQLPlayGround(context *gin.Context) {
	h := playground.Handler("GraphQL", "/query")

	h.ServeHTTP(context.Writer, context.Request)
}

func NewGraphQLController(
	resolver graphql.ResolverRoot,
) (*GraphQLController, error) {
	return &GraphQLController{
		resolver: resolver,
	}, nil
}
