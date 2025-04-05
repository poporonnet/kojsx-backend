package user

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/utils/password/argon2"
	"github.com/stretchr/testify/assert"
)

func TestLoginService_Login(t *testing.T) {
	d, _ := domain.NewUser("123", "test", "me@example.jp")
	d.SetVerified()
	enc := argon2.NewArgon2PasswordEncoder()
	encd, _ := enc.EncodePassword("hello")
	d.SetPassword(string(encd))

	r := inmemory.NewUserRepository([]domain.User{*d})
	s := NewLoginService(r, "")

	_, _, err := s.Login("me@example.jp", "hello")
	assert.Equal(t, nil, err)
}
