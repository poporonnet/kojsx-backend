package submission_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/application/submission"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestFindSubmissionService_FindByID(t *testing.T) {
	r := inmemory.NewSubmissionRepository(seed.NewSeeds().Submission)
	p := inmemory.NewProblemRepository(seed.NewSeeds().Problems)
	s := submission.NewFindSubmissionService(r, p)

	res, _ := s.FindByID("200")
	assert.Equal(t, submission.DomainToData(seed.NewSeeds().Submission[0]), res)
}
