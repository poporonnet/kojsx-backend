package service

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestContestantService_IsExists(t *testing.T) {
	r := inmemory.NewContestantRepository(dummyData.ContestantArray)
	s := NewContestantService(r)

	// trueのとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsContestantData))
	// falseのとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsContestantData))
}
