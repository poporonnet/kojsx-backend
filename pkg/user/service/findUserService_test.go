package service_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/user/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/stretchr/testify/assert"
)

func TestFindUserService_FindByID(t *testing.T) {
	r := inmemory.NewUserRepository(seed.NewSeeds().Users)
	s := service.NewFindUserService(r)

	// 取得できるとき
	res1, _ := s.FindByID("20")
	assert.Equal(t, service.DomainToData(seed.NewSeeds().Users[0]), res1)
	// 取得できないとき
	res2, _ := s.FindByID("0")
	assert.Equal(t, service.Data{}, res2)
}
