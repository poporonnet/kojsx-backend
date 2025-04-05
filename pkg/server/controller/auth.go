package controller

import (
	"fmt"

	"github.com/poporonnet/kojsx-backend/pkg/application/user"
	"github.com/poporonnet/kojsx-backend/pkg/repository"
	"github.com/poporonnet/kojsx-backend/pkg/server/controller/model"
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
		return model.LoginResponseJSON{}, fmt.Errorf("failed to login: %w", err)
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
