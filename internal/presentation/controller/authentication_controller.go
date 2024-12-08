package controller

import (
	"net/http"
	"original-card-game-backend/internal/application/usecase"
	"original-card-game-backend/internal/presentation/middleware"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationUsecase *usecase.AuthenticationUsecase
}

func (c *AuthenticationController) SignUp(context *gin.Context) {
	token, err := c.authenticationUsecase.SignUp()
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, map[string]string{
			"hoge": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, map[string]string{
		"token": token,
	})
}

func (c *AuthenticationController) Refresh(context *gin.Context) {
	tokenString := middleware.GetToken(context)

	token, err := c.authenticationUsecase.Refresh(tokenString)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, map[string]string{
			"hoge": err.Error(),
		})
		return
	}

	context.JSON(http.StatusCreated, map[string]string{
		"token": token,
	})
}

func NewAuthenticationController(
	authenticationUsecase *usecase.AuthenticationUsecase,
) (*AuthenticationController, error) {
	return &AuthenticationController{
		authenticationUsecase: authenticationUsecase,
	}, nil
}
