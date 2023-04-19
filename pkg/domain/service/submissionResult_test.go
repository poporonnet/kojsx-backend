package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubmissionResultService_IsExists(t *testing.T) {
	r := inmemory.NewSubmissionRepository(nil, dummyData.SubmissionResultArray)
	s := NewSubmissionResultService(r)
	// trueのとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsSubmissionResult))
	// falseのとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsSubmissionResult))
}
