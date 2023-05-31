package mongodb

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type UserRepository struct {
}

func (u UserRepository) CreateUser(d domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindAllUsers() []domain.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindUserByID(id id.SnowFlakeID) *domain.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindUserByName(name string) *domain.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) FindUserByEmail(email string) *domain.User {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) UpdateUser(d domain.User) error {
	//TODO implement me
	panic("implement me")
}
