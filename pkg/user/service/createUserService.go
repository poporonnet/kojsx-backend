package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/domainService"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
	"github.com/poporonnet/kojsx-backend/pkg/utils/mail"
	"github.com/poporonnet/kojsx-backend/pkg/utils/password/argon2"
	"github.com/poporonnet/kojsx-backend/pkg/utils/token"
)

const (
	verificationMailBody    = "KOJSへようこそ。\nKOJSにログインするには、以下のリンクをクリックしてください。\nメールアドレスを確認するとアカウントが有効化され、ログインができるようになります。\nリンクの有効期限は送信から24時間です。\nhttps://%s/verify/%s\n"
	verificationMailSubject = "Welcome to KOJS! (%s)"
)

type CreateUserService struct {
	userRepository repository.UserRepository
	userService    domainService.UserService
	idGenerator    id.Generator
	mailer         mail.Mailer
	key            string
}

func NewCreateUserService(userRepository repository.UserRepository, service domainService.UserService, mailer mail.Mailer, key string) *CreateUserService {
	return &CreateUserService{
		userRepository: userRepository,
		userService:    service,
		idGenerator:    id.NewSnowFlakeIDGenerator(),
		mailer:         mailer,
	}
}

func (s *CreateUserService) Handle(name string, password string, email string) (*model.User, string, error) {
	newID := s.idGenerator.NewID(time.Now())

	encoder := argon2.NewArgon2PasswordEncoder()
	pw, err := encoder.EncodePassword(password)
	if err != nil {
		return nil, "", err
	}

	d, err := model.NewUser(newID, name, email)
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
		return fmt.Errorf("failed to parse token: %w", err)
	}

	if tt.Type != "verify" {
		return errors.New("invalid token")
	}

	if tt.ID != id {
		return errors.New("user mismatched")
	}

	u, err := s.userRepository.FindUserByID(id)
	if err != nil {
		return fmt.Errorf("not found: %w", err)
	}

	u.SetVerified()

	err = s.userRepository.UpdateUser(*u)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
