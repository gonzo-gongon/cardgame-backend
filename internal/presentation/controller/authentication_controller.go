package controller

import (
	"net/http"
	"original-card-game-backend/internal/application/usecase"
	"original-card-game-backend/internal/presentation/middleware"
	"original-card-game-backend/internal/presentation/presenter"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	authenticationUsecase   *usecase.AuthenticationUsecase
	authenticationPresenter *presenter.AuthenticationPresenter
}

func (c *AuthenticationController) SignUp(ctx *gin.Context) {
	token, err := c.authenticationUsecase.SignUp()
	if err != nil {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			c.authenticationPresenter.Error(err),
		)

		return
	}

	ctx.JSON(
		http.StatusCreated,
		c.authenticationPresenter.Success(token),
	)
}

func (c *AuthenticationController) Refresh(ctx *gin.Context) {
	tokenString := middleware.GetToken(ctx)

	token, err := c.authenticationUsecase.Refresh(tokenString)
	if err != nil {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			c.authenticationPresenter.Error(err),
		)

		return
	}

	ctx.JSON(
		http.StatusCreated,
		c.authenticationPresenter.Success(token),
	)
}

func NewAuthenticationController(
	authenticationUsecase *usecase.AuthenticationUsecase,
	authenticationPresenter *presenter.AuthenticationPresenter,
) (*AuthenticationController, error) {
	return &AuthenticationController{
		authenticationUsecase:   authenticationUsecase,
		authenticationPresenter: authenticationPresenter,
	}, nil
}
