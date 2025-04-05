package controller_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/controller"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/controller/model"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/service"
	service2 "github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/utils/mail/dummy"
	"github.com/stretchr/testify/assert"
)

func TestUserController_Create(t *testing.T) {
	r := inmemory.NewUserRepository(seed.NewSeeds().Users)
	u := service.NewUserService(r)
	s := service2.NewCreateUserService(r, *u, dummy.NewMailer(), "")
	f := service2.NewFindUserService(r)
	c := controller.NewUserController(r, *s, *f)

	res, _ := c.Create(model.CreateUserRequestJSON{
		Name:     "miyoshi",
		Email:    "me@example.jp",
		Password: "hello",
	})

	assert.Equal(t, model.CreateUserResponseJSON{
		ID:    res.ID,
		Name:  "miyoshi",
		Email: "me@example.jp",
	}, res)
}
