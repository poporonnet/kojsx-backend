package submission_test

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/application/submission"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestFindSubmissionService_FindByID(t *testing.T) {
	r := inmemory.NewSubmissionRepository(seed.NewSeeds().Submission)
	s := submission.NewFindSubmissionService(r)

	res, _ := s.FindByID("200")
	assert.Equal(t, submission.DomainToData(seed.NewSeeds().Submission[0]), res)
}
