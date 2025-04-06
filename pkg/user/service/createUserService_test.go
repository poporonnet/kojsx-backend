package service_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/domainService"
	"github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/utils/mail/dummy"
	"github.com/stretchr/testify/assert"
)

func TestCreateUserService_Handle(t *testing.T) {
	r := inmemory.NewUserRepository(seed.NewSeeds().Users)
	uService := domainService.NewUserService(r)
	cUserService := service.NewCreateUserService(r, *uService, dummy.NewMailer(), "123")

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
	uService := domainService.NewUserService(r)
	cService := service.NewCreateUserService(r, *uService, dummy.NewMailer(), "123")

	d, token, _ := cService.Handle("miyoshi", "hello", "miyoshi@example.jp")
	err := cService.Verify(d.GetID(), token)
	assert.Equal(t, nil, err)
}
