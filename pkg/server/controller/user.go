package controller

import (
	"github.com/mct-joken/kojs5-backend/pkg/application/user"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
)

type UserController struct {
	repository    repository.UserRepository
	createService user.CreateUserService
}

func NewUserController(repository repository.UserRepository, createService user.CreateUserService) *UserController {
	return &UserController{repository: repository, createService: createService}
}

func (c *UserController) Create(req model.CreateUserRequestJSON) (model.CreateUserResponseJSON, error) {
	d, _, err := c.createService.Handle(req.Name, req.Password, req.Email)
	if err != nil {
		return model.CreateUserResponseJSON{}, err
	}

	return model.CreateUserResponseJSON{ID: string(d.GetID()), Name: d.GetName(), Email: d.GetEmail()}, nil
}
