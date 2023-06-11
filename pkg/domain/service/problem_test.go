package service

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestProblemService_IsExists(t *testing.T) {
	r := inmemory.NewProblemRepository(dummyData.ProblemArray)
	s := NewProblemService(r)

	// trueになるとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsProblem))

	// falseになるとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsProblem))
}
