package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProblemService_IsExists(t *testing.T) {
	r := inmemory.NewProblemRepository(dummyData.ProblemArray, nil, nil)
	s := NewProblemService(r)

	// trueになるとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsProblem))

	// falseになるとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsProblem))
	
}
