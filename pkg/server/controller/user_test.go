package controller_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/server/controller"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/application/user"
	"github.com/poporonnet/kojsx-backend/pkg/domain/service"
	"github.com/poporonnet/kojsx-backend/pkg/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/server/controller/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/mail/dummy"
	"github.com/stretchr/testify/assert"
)

func TestUserController_Create(t *testing.T) {
	r := inmemory.NewUserRepository(seed.NewSeeds().Users)
	u := service.NewUserService(r)
	s := user.NewCreateUserService(r, *u, dummy.NewMailer(), "")
	f := user.NewFindUserService(r)
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
