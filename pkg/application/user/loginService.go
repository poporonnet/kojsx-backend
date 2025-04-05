package user

import (
	"errors"
	"fmt"

	"github.com/poporonnet/kojsx-backend/pkg/repository"
	password2 "github.com/poporonnet/kojsx-backend/pkg/utils/password"
	"github.com/poporonnet/kojsx-backend/pkg/utils/password/argon2"
	"github.com/poporonnet/kojsx-backend/pkg/utils/token"
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
		return "", "", fmt.Errorf("failed to find user: %w", err)
	}

	if !res.IsVerified() {
		return "", "", errors.New("not verified")
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

func (s *LoginService) Verify(t string) bool {
	parser := token.NewJWTTokenParser(s.key)
	d, err := parser.Parse(t)
	if err != nil {
		return false
	}
	if d.Type != "verify" {
		return false
	}

	u, err := s.findService.FindByID(d.ID)
	if err != nil {
		return false
	}
	do := u.ToDomain()
	do.SetVerified()
	// ToDo: updateService作る
	err = s.repository.UpdateUser(do)

	return err == nil
}
