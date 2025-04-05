package contest_test

import (
	"errors"
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/application/contest"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/domain/service"
	"github.com/poporonnet/kojsx-backend/pkg/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/utils/dummyData"
	"github.com/stretchr/testify/assert"
)

func TestJoinContestService_Join(t *testing.T) {
	r := inmemory.NewContestantRepository(dummyData.ContestantArray)
	s := contest.NewJoinContestService(r, *service.NewContestantService(r))

	// 作れるとき
	err := s.Join("10", seed.NewSeeds().Users[0], 0)
	assert.Equal(t, nil, err)

	// 作れないとき
	err2 := s.Join("10", seed.NewSeeds().Users[0], 0)
	assert.Equal(t, errors.New("AlreadyJoined"), err2)
}
