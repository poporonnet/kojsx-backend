package controller

import (
	"fmt"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/controller/model"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/user/service"
)

type AuthController struct {
	repository   repository.UserRepository
	loginService service.LoginService
}

func NewAuthController(userRepository repository.UserRepository, key string) *AuthController {
	return &AuthController{repository: userRepository, loginService: *service.NewLoginService(userRepository, key)}
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
