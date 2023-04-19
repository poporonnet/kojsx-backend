package user

import (
	"errors"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type CreateUserService struct {
	userRepository repository.UserRepository
	userService    service.UserService
	idGenerator    id.Generator
}

func NewCreateUserService(userRepository repository.UserRepository, service service.UserService) *CreateUserService {
	return &CreateUserService{
		userRepository: userRepository,
		userService:    service,
		idGenerator:    id.NewSnowFlakeIDGenerator(),
	}
}

func (s *CreateUserService) Handle(name string, email string) error {
	newID := s.idGenerator.NewID(time.Now())
	d, err := domain.NewUser(newID, name, email)
	if err != nil {
		return err
	}
	exists := s.userService.IsExists(*d)
	if exists {
		return errors.New("UserExists")
	}
	res := s.userRepository.CreateUser(*d)
	if res != nil {
		return res
	}
	return nil
}
