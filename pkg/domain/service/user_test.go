package service_test

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestUserService_IsExists(t *testing.T) {
	// trueになる場合
	repo := inmemory.NewUserRepository(seed.NewSeeds().Users)
	s := service.NewUserService(repo)

	assert.Equal(t, true, s.IsExists(seed.NewSeeds().Users[0]))

	// falseになる場合
	assert.Equal(t, false, s.IsExists(*dummyData.NotExists))
}
