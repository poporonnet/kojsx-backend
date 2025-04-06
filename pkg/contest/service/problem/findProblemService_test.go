package problem_test

import (
	"testing"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/problem"
	"github.com/stretchr/testify/assert"

	"github.com/poporonnet/kojsx-backend/pkg/utils"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"
)

var pService problem.FindProblemService

func TestMain(t *testing.M) {
	utils.NewLogger()
	pService = *problem.NewFindProblemService(
		inmemory.NewProblemRepository(seed.NewSeeds().Problems),
		inmemory.NewContestRepository(seed.NewSeeds().Contests),
		inmemory.NewContestantRepository(seed.NewSeeds().Contestants),
	)
	t.Run()
}

func TestFindProblemService_FindByID(t *testing.T) {
	act := []*problem.Data{
		func() *problem.Data {
			r, _ := pService.FindByID("110", time.Now(), "20")
			return r
		}(),
		func() *problem.Data {
			r, _ := pService.FindByID("110", time.Date(2020, time.April, 1, 12, 0, 0, 0, time.UTC), "30")
			return r
		}(),
	}

	t.Run("コンテスト開始後は取得できる", func(t *testing.T) {
		assert.Equal(t, problem.DomainToData(seed.NewSeeds().Problems[0]), *act[0])
	})
	t.Run("コンテスト開始前は取得できない", func(t *testing.T) {
		assert.Equal(t, (*problem.Data)(nil), act[1])
	})
	t.Run("コンテスト開始前でもコンテスト管理者, テスターは取得できる", func(t *testing.T) {
		assert.NotEqual(t, (*problem.Data)(nil), *act[0])
	})
}

func TestFindProblemService_FindByContestID(t *testing.T) {

}
