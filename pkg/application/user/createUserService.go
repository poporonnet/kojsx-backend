package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"github.com/mct-joken/kojs5-backend/pkg/utils/mail"
	"github.com/mct-joken/kojs5-backend/pkg/utils/password/argon2"
	"github.com/mct-joken/kojs5-backend/pkg/utils/token"
)

const (
	verificationMailBody    = "KOJSへようこそ。\nKOJSにログインするには、以下のリンクをクリックしてください。\nメールアドレスを確認するとアカウントが有効化され、ログインができるようになります。\nリンクの有効期限は送信から24時間です。\nhttps://%s/verify/%s\n"
	verificationMailSubject = "Welcome to KOJS! (%s)"
)

type CreateUserService struct {
	userRepository repository.UserRepository
	userService    service.UserService
	idGenerator    id.Generator
	mailer         mail.Mailer
	key            string
}

func NewCreateUserService(userRepository repository.UserRepository, service service.UserService, mailer mail.Mailer, key string) *CreateUserService {
	return &CreateUserService{
		userRepository: userRepository,
		userService:    service,
		idGenerator:    id.NewSnowFlakeIDGenerator(),
		mailer:         mailer,
	}
}

func (s *CreateUserService) Handle(name string, password string, email string) (*domain.User, string, error) {
	newID := s.idGenerator.NewID(time.Now())

	encoder := argon2.NewArgon2PasswordEncoder()
	pw, err := encoder.EncodePassword(password)
	if err != nil {
		return nil, "", err
	}

	d, err := domain.NewUser(newID, name, email)
	if err != nil {
		return nil, "", err
	}

	d.SetPassword(string(pw))

	exists := s.userService.IsExists(*d)
	if exists {
		return nil, "", errors.New("UserExists")
	}
	res := s.userRepository.CreateUser(*d)
	if res != nil {
		return nil, "", res
	}

	t := token.NewJWTTokenGenerator(s.key)
	tt, err := t.NewVerifyToken(d.GetID())
	if err != nil {
		return nil, "", err
	}

	// ToDo: メール本文に書くリンクを設定できるようにする
	err = s.mailer.Send(email,
		fmt.Sprintf(verificationMailBody, "ojs.kosen.dev", tt),
		fmt.Sprintf(verificationMailSubject, "ojs.kosen.dev"))
	if err != nil {
		return nil, "", err
	}
	return d, tt, nil
}

func (s *CreateUserService) Verify(id id.SnowFlakeID, t string) error {
	p := token.NewJWTTokenParser(s.key)
	tt, err := p.Parse(t)
	if err != nil {
		return err
	}

	if tt.Type != "verify" {
		return errors.New("invalid token")
	}

	if tt.ID != id {
		return errors.New("user mismatched")
	}

	u := s.userRepository.FindUserByID(id)
	if u == nil {
		return errors.New("no such user")
	}

	u.SetVerified()

	err = s.userRepository.UpdateUser(*u)
	if err != nil {
		return err
	}

	return nil
}
