package server

import (
	"fmt"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/controller/schema"
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

func (c *AuthController) Login(req schema.LoginRequestJSON) (schema.LoginResponseJSON, error) {
	a, r, err := c.loginService.Login(req.Email, req.Password)
	if err != nil {
		return schema.LoginResponseJSON{}, fmt.Errorf("failed to login: %w", err)
	}

	return schema.LoginResponseJSON{
		AccessToken:  a,
		RefreshToken: r,
	}, nil
}

func (c *AuthController) Verify(token string) (bool, error) {
	res := c.loginService.Verify(token)
	return res, nil
}
