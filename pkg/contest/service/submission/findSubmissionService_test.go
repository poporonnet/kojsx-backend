package submission_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/repository/inmemory"
	submission2 "github.com/poporonnet/kojsx-backend/pkg/contest/service/submission"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/stretchr/testify/assert"
)

func TestFindSubmissionService_FindByID(t *testing.T) {
	r := inmemory.NewSubmissionRepository(seed.NewSeeds().Submission)
	p := inmemory.NewProblemRepository(seed.NewSeeds().Problems)
	s := submission2.NewFindSubmissionService(r, p)

	res, _ := s.FindByID("200")
	assert.Equal(t, submission2.DomainToData(seed.NewSeeds().Submission[0]), res)
}
