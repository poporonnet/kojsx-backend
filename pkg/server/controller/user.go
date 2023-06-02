package controller

import (
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/application/user"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
)

type UserController struct {
	repository    repository.UserRepository
	createService user.CreateUserService
	findService   user.FindUserService
}

func NewUserController(repository repository.UserRepository, createService user.CreateUserService, service user.FindUserService) *UserController {
	return &UserController{repository: repository, createService: createService, findService: service}
}

func (c *UserController) Create(req model.CreateUserRequestJSON) (model.CreateUserResponseJSON, error) {
	d, _, err := c.createService.Handle(req.Name, req.Password, req.Email)
	if err != nil {
		return model.CreateUserResponseJSON{}, fmt.Errorf("failed to create user: %w", err)
	}

	return model.CreateUserResponseJSON{ID: string(d.GetID()), Name: d.GetName(), Email: d.GetEmail()}, nil
}

func (c *UserController) FindAllUsers() ([]model.FindUsersResponseJSON, error) {
	d, err := c.findService.FindAllUsers()
	if err != nil {
		return []model.FindUsersResponseJSON{}, fmt.Errorf("failed to find users: %w", err)
	}

	res := make([]model.FindUsersResponseJSON, len(d))
	for i, v := range d {
		role := 0
		if !v.IsAdmin() {
			role = 1
		}

		res[i] = model.FindUsersResponseJSON{
			ID:   string(v.GetID()),
			Name: v.GetName(),
			Role: role,
		}
	}

	return res, nil
}
