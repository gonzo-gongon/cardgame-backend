package app

import (
	// "backend/internal/presentation/controller"

	"github.com/gin-gonic/gin"
)

// main関数はunused対象外とする
func main() { //nolint:unused
	r := gin.Default()

	// panicが出るのでエラーチェックしない
	r.Run(":8080") //nolint:errcheck
}
