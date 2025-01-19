package controller

import (
	"original-card-game-backend/internal/presentation/graphql"
	"original-card-game-backend/internal/presentation/graphql/directive"
	graphqlgen "original-card-game-backend/internal/presentation/graphql/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

type GraphQLController struct {
	resolver *graphql.Resolver
}

func (c *GraphQLController) GraphQL(context *gin.Context, ad *directive.AuthDirective) {
	h := handler.New(
		graphqlgen.NewExecutableSchema(
			graphqlgen.Config{
				Resolvers: c.resolver,
				Directives: graphqlgen.DirectiveRoot{
					Auth: ad.Auth,
				},
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
	resolver *graphql.Resolver,
) (*GraphQLController, error) {
	return &GraphQLController{
		resolver: resolver,
	}, nil
}
