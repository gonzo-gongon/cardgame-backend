package controller

import (
	"net/http"
	"original-card-game-backend/internal/application/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase *usecase.UserUsecase
}

func (c *UserController) GetOne(context *gin.Context) {
	_, err := c.userUsecase.GetUser("0193a617-95b2-7451-82f5-1059ef445867")

	if err != nil {
		context.JSON(http.StatusNotFound, map[string]string{
			"hoge": "NotFound",
		})
		return
	}

	context.JSON(http.StatusOK, map[string]string{
		"hoge": "OK",
	})
}

func NewUserController(
	userUsecase *usecase.UserUsecase,
) (*UserController, error) {
	return &UserController{
		userUsecase: userUsecase,
	}, nil
}
