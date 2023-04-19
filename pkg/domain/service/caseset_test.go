package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCaseSetService_IsExists(t *testing.T) {
	r := inmemory.NewProblemRepository(nil, dummyData.CasesetArray, nil)
	s := NewCaseSetService(r)

	// trueのとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsCasesetData))
	// falseのとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsCasesetData))
}
