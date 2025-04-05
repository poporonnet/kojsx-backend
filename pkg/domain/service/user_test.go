package service_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/domain/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/utils/dummyData"
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
