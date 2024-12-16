package main

import (
	// "backend/internal/presentation/controller"

	"original-card-game-backend/cmd/app"
	"original-card-game-backend/internal/presentation/controller"
	"original-card-game-backend/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

// main関数はunused対象外とする
func main() { //nolint:unused
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

	// panicが出るのでエラーチェックしない
	server.Run(":8080") //nolint:errcheck
}
