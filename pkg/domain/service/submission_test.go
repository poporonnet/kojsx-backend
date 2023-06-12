package service_test

import (
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"

	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestSubmissionService_IsExists(t *testing.T) {
	r := inmemory.NewSubmissionRepository(seed.NewSeeds().Submission)
	s := service.NewSubmissionService(r)
	// trueのとき
	assert.Equal(t, true, s.IsExists(seed.NewSeeds().Submission[0]))
	// falseのとき
	assert.Equal(t, false, s.IsExists(*dummyData.NotExistsSubmission))
}
