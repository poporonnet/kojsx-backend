package inmemory

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type UserRepository struct {
	data []domain.User
}

func NewUserRepository(d []domain.User) *UserRepository {
	return &UserRepository{data: d}
}

func (u UserRepository) FindUserByID(id id.SnowFlakeID) *domain.User {
	for _, v := range u.data {
		if v.GetID() == id {
			return &v
		}
	}
	return nil
}

func (u UserRepository) FindUserByName(name string) *domain.User {
	for _, v := range u.data {
		if v.GetName() == name {
			return &v
		}
	}
	return nil
}

func (u UserRepository) FindUserByEmail(email string) *domain.User {
	for _, v := range u.data {
		if v.GetEmail() == email {
			return &v
		}
	}
	return nil
}
