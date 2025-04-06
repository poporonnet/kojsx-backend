package controller_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/controller"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/controller/schema"
	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/user/model/domainService"
	"github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/utils/mail/dummy"
	"github.com/stretchr/testify/assert"
)

func TestUserController_Create(t *testing.T) {
	r := inmemory.NewUserRepository(seed.NewSeeds().Users)
	u := domainService.NewUserService(r)
	s := service.NewCreateUserService(r, *u, dummy.NewMailer(), "")
	f := service.NewFindUserService(r)
	c := controller.NewUserController(r, *s, *f)

	res, _ := c.Create(schema.CreateUserRequestJSON{
		Name:     "miyoshi",
		Email:    "me@example.jp",
		Password: "hello",
	})

	assert.Equal(t, schema.CreateUserResponseJSON{
		ID:    res.ID,
		Name:  "miyoshi",
		Email: "me@example.jp",
	}, res)
}
