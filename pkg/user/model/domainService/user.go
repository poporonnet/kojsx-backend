package domainService

import (
	"github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{userRepository: repository}
}

func (s *UserService) IsExists(u model.User) bool {
	// 重複判定: ユーザー名/ID/Email
	i, _ := s.userRepository.FindUserByID(u.GetID())
	n, _ := s.userRepository.FindUserByName(u.GetName())
	e, _ := s.userRepository.FindUserByEmail(u.GetEmail())
	if i == nil && n == nil && e == nil {
		return false
	}

	return true
}
