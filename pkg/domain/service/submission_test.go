package service

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestSubmissionService_IsExists(t *testing.T) {
	r := inmemory.NewSubmissionRepository(dummyData.SubmissionArray, nil)
	s := NewSubmissionService(r)
	// trueのとき
	assert.Equal(t, true, s.IsExists(*dummyData.ExistsSubmission))
	// falseのとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsSubmission))
}
