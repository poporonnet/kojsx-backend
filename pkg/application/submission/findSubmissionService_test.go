package submission

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestFindSubmissionService_FindByID(t *testing.T) {
	r := inmemory.NewSubmissionRepository(dummyData.SubmissionArray, dummyData.SubmissionResultArray)
	s := NewFindSubmissionService(r)

	res, _ := s.FindByID("1")
	assert.Equal(t, dummyData.ExistsSubmission, res)
}
