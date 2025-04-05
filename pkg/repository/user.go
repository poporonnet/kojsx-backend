package repository

import (
	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type UserRepository interface {
	CreateUser(d domain.User) error
	FindAllUsers() ([]domain.User, error)
	FindUserByID(id id.SnowFlakeID) (*domain.User, error)
	FindUserByName(name string) (*domain.User, error)
	FindUserByEmail(email string) (*domain.User, error)
	UpdateUser(d domain.User) error
}
