package main

import (
	// "backend/internal/presentation/controller"

	"original-card-game-backend/cmd/app"
	"original-card-game-backend/internal/presentation/controller"

	"github.com/gin-gonic/gin"
)

// main関数はunused対象外とする
func main() { //nolint:unused
	container := app.BuildContainer()

	r := gin.Default()

	if err := container.Invoke(func(c *controller.UserController) {
		r.GET("/user", c.GetOne)
	}); err != nil {
		panic(err)
	}

	// panicが出るのでエラーチェックしない
	r.Run(":8080") //nolint:errcheck
}
