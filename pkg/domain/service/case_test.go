package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaseService_IsExists(t *testing.T) {
	r := inmemory.NewProblemRepository(nil, nil, dummyData.CaseArray)
	s := NewCaseService(r)

	// trueのとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsCase))
	// falseのとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsCase))
}
