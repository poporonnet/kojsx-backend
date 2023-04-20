package user

import (
	"errors"

	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type FindUserService struct {
	userRepository repository.UserRepository
}

func NewFindUserService(userRepository repository.UserRepository) *FindUserService {
	return &FindUserService{userRepository: userRepository}
}

func (s *FindUserService) FindByID(id id.SnowFlakeID) (Data, error) {
	u := s.userRepository.FindUserByID(id)
	if u == nil {
		return Data{}, errors.New("")
	}
	return DomainToData(*u), nil
}
