package controller

import (
	"fmt"

	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/controller/model"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type UserController struct {
	repository    repository.UserRepository
	createService service.CreateUserService
	findService   service.FindUserService
}

func NewUserController(repository repository.UserRepository, createService service.CreateUserService, service service.FindUserService) *UserController {
	return &UserController{repository: repository, createService: createService, findService: service}
}

func (c *UserController) Create(req model.CreateUserRequestJSON) (model.CreateUserResponseJSON, error) {
	d, _, err := c.createService.Handle(req.Name, req.Password, req.Email)
	if err != nil {
		return model.CreateUserResponseJSON{}, fmt.Errorf("failed to create user: %w", err)
	}

	return model.CreateUserResponseJSON{ID: string(d.GetID()), Name: d.GetName(), Email: d.GetEmail()}, nil
}

func (c *UserController) FindByID(uID string) (model.FindUsersResponseJSON, error) {
	res, err := c.findService.FindByID(id.SnowFlakeID(uID))
	if err != nil {
		return model.FindUsersResponseJSON{}, err
	}
	return model.FindUsersResponseJSON{
		ID:   string(res.GetID()),
		Name: res.GetName(),
		Role: func() int {
			if !res.IsVerified() {
				return 2
			}
			if res.IsAdmin() {
				return 0
			}
			return 1
		}(),
	}, err
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
