package presenter

import "original-card-game-backend/internal/presentation/response"

type AuthenticationPresenter struct{}

func (p *AuthenticationPresenter) Success(token string) response.Response {
	return response.Success(
		response.AuthenticationResponse{
			Token: token,
		},
	)
}

func (p *AuthenticationPresenter) Error(err error) response.Response {
	return response.Error(err.Error())
}

func NewAuthenticationPresenter() *AuthenticationPresenter {
	return &AuthenticationPresenter{}
}
