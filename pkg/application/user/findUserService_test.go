package user

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestFindUserService_FindByID(t *testing.T) {
	r := inmemory.NewUserRepository(dummyData.UserArray)
	s := NewFindUserService(r)

	// 取得できるとき
	res1, _ := s.FindByID("1")
	assert.Equal(t, DomainToData(dummyData.Exists), res1)
	// 取得できないとき
	res2, _ := s.FindByID("0")
	assert.Equal(t, Data{}, res2)
}
