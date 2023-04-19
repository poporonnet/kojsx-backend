package repository

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type UserRepository interface {
	CreateUser(d domain.User) error
	FindUserByID(id id.SnowFlakeID) *domain.User
	FindUserByName(name string) *domain.User
	FindUserByEmail(email string) *domain.User
}
