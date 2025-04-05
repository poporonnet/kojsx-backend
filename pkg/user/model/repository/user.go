package repository

import (
	"github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type UserRepository interface {
	CreateUser(d model.User) error
	FindAllUsers() ([]model.User, error)
	FindUserByID(id id.SnowFlakeID) (*model.User, error)
	FindUserByName(name string) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
	UpdateUser(d model.User) error
}
