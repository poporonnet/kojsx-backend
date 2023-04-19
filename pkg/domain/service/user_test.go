package service

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestUserService_IsExists(t *testing.T) {
	// trueになる場合
	repo := inmemory.NewUserRepository(dummyData.UserArray)
	s := NewUserService(repo)

	assert.Equal(t, true, s.IsExists(dummyData.Exists))

	// falseになる場合
	assert.Equal(t, false, s.IsExists(*dummyData.NotExists))
}
