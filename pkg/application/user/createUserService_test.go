package user_test

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/application/user"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"

	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/mail/dummy"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserService_Handle(t *testing.T) {
	r := inmemory.NewUserRepository(seed.NewSeeds().Users)
	uService := service.NewUserService(r)
	cUserService := user.NewCreateUserService(r, *uService, dummy.NewMailer(), "123")

	// 成功するとき
	_, _, err := cUserService.Handle("miyoshi", "hello", "miyoshi@example.jp")
	assert.Equal(t, nil, err)

	// 失敗するとき
	// 重複したとき
	_, _, err2 := cUserService.Handle("miyoshi", "hello", "miyoshi@example.jp")
	assert.NotEqual(t, nil, err2)
}

func TestCreateUserService_Verify(t *testing.T) {
	r := inmemory.NewUserRepository(seed.NewSeeds().Users)
	uService := service.NewUserService(r)
	cService := user.NewCreateUserService(r, *uService, dummy.NewMailer(), "123")

	d, token, _ := cService.Handle("miyoshi", "hello", "miyoshi@example.jp")
	err := cService.Verify(d.GetID(), token)
	assert.Equal(t, nil, err)
}
