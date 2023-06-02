package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

func (s *UserService) IsExists(u domain.User) bool {
	// 重複判定: ユーザー名/ID/Email
	i, _ := s.userRepository.FindUserByID(u.GetID())
	n, _ := s.userRepository.FindUserByName(u.GetName())
	e, _ := s.userRepository.FindUserByEmail(u.GetEmail())
	if i == nil && n == nil && e == nil {
		return false
	}

	return true
}
