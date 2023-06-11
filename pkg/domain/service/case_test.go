package service

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestCaseService_IsExists(t *testing.T) {
	r := inmemory.NewProblemRepository(dummyData.ProblemArray)
	s := NewCaseService(r)

	// trueのとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsCase))
	// falseのとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsCase))
}
