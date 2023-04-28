package controller

import (
	"github.com/mct-joken/kojs5-backend/pkg/application/user"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
)

type AuthController struct {
	repository   repository.UserRepository
	loginService user.LoginService
}

func NewAuthController(userRepository repository.UserRepository, key string) *AuthController {
	return &AuthController{repository: userRepository, loginService: *user.NewLoginService(userRepository, key)}
}

func (c *AuthController) Login(req model.LoginRequestJSON) (model.LoginResponseJSON, error) {
	a, r, err := c.loginService.Login(req.Email, req.Password)
	if err != nil {
		return model.LoginResponseJSON{}, err
	}

	return model.LoginResponseJSON{
		AccessToken:  a,
		RefreshToken: r,
	}, nil
}

func (c *AuthController) Verify(token string) (bool, error) {
	res := c.loginService.Verify(token)
	return res, nil
}
