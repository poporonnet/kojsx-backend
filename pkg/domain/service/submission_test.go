package service_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/domain/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/utils/dummyData"
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
