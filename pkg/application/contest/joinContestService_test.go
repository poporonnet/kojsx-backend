package contest_test

import (
	"errors"
	"github.com/mct-joken/kojs5-backend/pkg/application/contest"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"
	"testing"

	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils/dummyData"
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
