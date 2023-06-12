package user_test

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/application/user"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestFindUserService_FindByID(t *testing.T) {
	r := inmemory.NewUserRepository(seed.NewSeeds().Users)
	s := user.NewFindUserService(r)

	// 取得できるとき
	res1, _ := s.FindByID("20")
	assert.Equal(t, user.DomainToData(seed.NewSeeds().Users[0]), res1)
	// 取得できないとき
	res2, _ := s.FindByID("0")
	assert.Equal(t, user.Data{}, res2)
}
