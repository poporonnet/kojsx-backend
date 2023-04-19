package user

import (
	"errors"

	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
)

type CreateUserService struct {
	userRepository repository.UserRepository
	userService    service.UserService
}

func NewCreateUserService(userRepository repository.UserRepository, service service.UserService) *CreateUserService {
	return &CreateUserService{
		userRepository: userRepository,
		userService:    service,
	}
}

func (s *CreateUserService) Handle(name string, email string) error {
	return errors.New("NotImplemented")
}
