package service_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestContestService_IsExists(t *testing.T) {
	r := inmemory.NewContestRepository(seed.NewSeeds().Contests)
	s := service.NewContestService(r)
	// trueになるとき
	assert.Equal(t, true, s.IsExists(seed.NewSeeds().Contests[0]))
	// falseになるとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsContestData))
}
