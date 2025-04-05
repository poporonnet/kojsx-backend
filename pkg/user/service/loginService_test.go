package service

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/password/argon2"
	"github.com/stretchr/testify/assert"
)

func TestLoginService_Login(t *testing.T) {
	d, _ := model.NewUser("123", "test", "me@example.jp")
	d.SetVerified()
	enc := argon2.NewArgon2PasswordEncoder()
	encd, _ := enc.EncodePassword("hello")
	d.SetPassword(string(encd))

	r := inmemory.NewUserRepository([]model.User{*d})
	s := NewLoginService(r, "")

	_, _, err := s.Login("me@example.jp", "hello")
	assert.Equal(t, nil, err)
}
