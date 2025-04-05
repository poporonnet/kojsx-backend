package user

import (
	"fmt"

	"github.com/poporonnet/kojsx-backend/pkg/repository"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type FindUserService struct {
	userRepository repository.UserRepository
}

func NewFindUserService(userRepository repository.UserRepository) *FindUserService {
	return &FindUserService{userRepository: userRepository}
}

func (s *FindUserService) FindAllUsers() ([]Data, error) {
	r, err := s.userRepository.FindAllUsers()
	if err != nil {
		return nil, fmt.Errorf("failed to find all users: %w", err)
	}
	u := make([]Data, len(r))

	for i, v := range r {
		u[i] = DomainToData(v)
	}

	return u, nil
}

func (s *FindUserService) FindByID(id id.SnowFlakeID) (Data, error) {
	u, err := s.userRepository.FindUserByID(id)
	if err != nil {
		return Data{}, fmt.Errorf("failed to find user: %w", err)
	}
	return DomainToData(*u), nil
}

func (s *FindUserService) FindUserByEmail(email string) (Data, error) {
	u, err := s.userRepository.FindUserByEmail(email)
	if err != nil {
		return Data{}, fmt.Errorf("failed to find user: %w", err)
	}
	return DomainToData(*u), nil
}
