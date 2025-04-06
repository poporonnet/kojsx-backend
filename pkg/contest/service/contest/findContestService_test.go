package contest_test

import (
	"testing"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/contest"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/stretchr/testify/assert"
)

func TestFindContestService_FindByID(t *testing.T) {
	r := inmemory.NewContestRepository(seed.NewSeeds().Contests)
	s := contest.NewFindContestService(r)

	// 取得できる
	r1, _ := s.FindByID("10")
	assert.Equal(t, contest.DomainToData(seed.NewSeeds().Contests[0]), *r1)
	// 取得できない
	r2, _ := s.FindByID("9")
	var ex *contest.Data = nil
	assert.Equal(t, ex, r2)
}

func TestFindContestService_FindAll(t *testing.T) {
	r := inmemory.NewContestRepository(seed.NewSeeds().Contests)
	s := contest.NewFindContestService(r)

	// 取得できる
	r1, _ := s.FindAll()
	ex := make([]contest.Data, len(seed.NewSeeds().Contests))
	for i, v := range seed.NewSeeds().Contests {
		ex[i] = contest.DomainToData(v)
	}

	assert.Equal(t, ex, r1)
}
