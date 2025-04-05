package inmemory

import (
	"errors"

	"github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type UserRepository struct {
	data []model.User
}

func NewUserRepository(d []model.User) *UserRepository {
	return &UserRepository{data: d}
}

func (u *UserRepository) CreateUser(d model.User) error {
	u.data = append(u.data, d)
	return nil
}

func (u *UserRepository) FindAllUsers() ([]model.User, error) {
	return u.data, nil
}

func (u *UserRepository) FindUserByID(id id.SnowFlakeID) (*model.User, error) {
	for _, v := range u.data {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("no such user")
}

func (u *UserRepository) FindUserByName(name string) (*model.User, error) {
	for _, v := range u.data {
		if v.GetName() == name {
			return &v, nil
		}
	}
	return nil, errors.New("no such user")
}

func (u *UserRepository) FindUserByEmail(email string) (*model.User, error) {
	for _, v := range u.data {
		if v.GetEmail() == email {
			return &v, nil
		}
	}
	return nil, errors.New("no such user")
}

func (u *UserRepository) UpdateUser(d model.User) error {
	if _, e := u.FindUserByID(d.GetID()); e != nil {
		return errors.New("no such user")
	}

	for i, v := range u.data {
		if v.GetID() == d.GetID() {
			u.data[i] = d
		}
	}

	return nil
}
