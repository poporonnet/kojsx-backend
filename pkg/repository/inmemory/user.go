package inmemory

import (
	"errors"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type UserRepository struct {
	data []domain.User
}

func NewUserRepository(d []domain.User) *UserRepository {
	return &UserRepository{data: d}
}

func (u *UserRepository) CreateUser(d domain.User) error {
	u.data = append(u.data, d)
	return nil
}

func (u *UserRepository) FindUserByID(id id.SnowFlakeID) *domain.User {
	for _, v := range u.data {
		if v.GetID() == id {
			return &v
		}
	}
	return nil
}

func (u *UserRepository) FindUserByName(name string) *domain.User {
	for _, v := range u.data {
		if v.GetName() == name {
			return &v
		}
	}
	return nil
}

func (u *UserRepository) FindUserByEmail(email string) *domain.User {
	for _, v := range u.data {
		if v.GetEmail() == email {
			return &v
		}
	}
	return nil
}

func (u *UserRepository) UpdateUser(d domain.User) error {
	if d := u.FindUserByID(d.GetID()); d == nil {
		return errors.New("no such user")
	}

	for i, v := range u.data {
		if v.GetID() == d.GetID() {
			u.data[i] = d
		}
	}

	return nil
}
