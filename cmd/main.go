package main

import (
	// "backend/internal/presentation/controller"

	"original-card-game-backend/cmd/app"
	"original-card-game-backend/internal/presentation/controller"
	"original-card-game-backend/internal/presentation/graphql/directive"
	"original-card-game-backend/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

func main() { //nolint:unused // main関数はunused対象外とする
	container := app.BuildContainer()

	server := gin.Default()

	if err := container.Invoke(
		func(
			c *controller.AuthenticationController,
			m *middleware.TokenRefreshMiddleware,
		) {
			server.POST("/signup", c.SignUp)
			server.POST("/refresh", m.Handler(), c.Refresh)
		},
	); err != nil {
		panic(err)
	}

	if err := container.Invoke(
		func(
			c *controller.GraphQLController,
			m *middleware.AuthenticationMiddleware,
			cm *middleware.ContextMiddleware,
			ad *directive.AuthDirective,
		) {
			server.Use(cm.Bind())

			server.GET("/", c.GraphQLPlayGround)
			server.POST("/query", m.Handler(), func(ctx *gin.Context) {
				c.GraphQL(ctx, ad)
			})
		},
	); err != nil {
		panic(err)
	}

	server.Run(":8080") //nolint:errcheck // panicが出るのでエラーチェックしない
}
