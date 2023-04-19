package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContestService_IsExists(t *testing.T) {
	r := inmemory.NewContestRepository(dummyData.ContestArray)
	s := NewContestService(r)
	// trueになるとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsContestData))
	// falseになるとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsContestData))
}
