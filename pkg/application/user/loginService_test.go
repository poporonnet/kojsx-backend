package user

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/password/argon2"
	"github.com/stretchr/testify/assert"
)

func TestLoginService_Login(t *testing.T) {
	d, _ := domain.NewUser("123", "test", "me@example.jp")
	enc := argon2.NewArgon2PasswordEncoder()
	encd, _ := enc.EncodePassword("hello")
	d.SetPassword(string(encd))

	r := inmemory.NewUserRepository([]domain.User{*d})
	s := NewLoginService(r, "")

	_, _, err := s.Login("me@example.jp", "hello")
	assert.Equal(t, nil, err)
}
