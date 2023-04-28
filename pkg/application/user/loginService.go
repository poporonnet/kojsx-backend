package user

import (
	"errors"

	"github.com/mct-joken/kojs5-backend/pkg/repository"
	password2 "github.com/mct-joken/kojs5-backend/pkg/utils/password"
	"github.com/mct-joken/kojs5-backend/pkg/utils/password/argon2"
	"github.com/mct-joken/kojs5-backend/pkg/utils/token"
)

type LoginService struct {
	repository  repository.UserRepository
	findService FindUserService
	key         string
}

func NewLoginService(repository repository.UserRepository, key string) *LoginService {
	return &LoginService{
		repository:  repository,
		findService: *NewFindUserService(repository),
		key:         key,
	}
}

func (s *LoginService) Login(email string, password string) (string, string, error) {
	res, err := s.findService.FindUserByEmail(email)
	if err != nil {
		return "", "", err
	}
	enc := argon2.NewArgon2PasswordEncoder()
	if !enc.IsMatchPassword(password, password2.EncodedPassword(res.GetPassword())) {
		return "", "", errors.New("password not matched")
	}

	g := token.NewJWTTokenGenerator(s.key)
	access, _ := g.NewAccessToken(res.GetID())
	refresh, _ := g.NewRefreshToken(res.GetID())
	return access, refresh, nil
}
