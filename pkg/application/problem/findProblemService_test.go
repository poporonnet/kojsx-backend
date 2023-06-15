package problem_test

import (
	"github.com/mct-joken/kojs5-backend/pkg/application/problem"
	"github.com/mct-joken/kojs5-backend/pkg/repository/inmemory"
	"github.com/mct-joken/kojs5-backend/pkg/utils"
	"github.com/mct-joken/kojs5-backend/pkg/utils/seed"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var pService problem.FindProblemService

func TestMain(t *testing.M) {
	utils.NewLogger()
	pService = *problem.NewFindProblemService(
		inmemory.NewProblemRepository(seed.NewSeeds().Problems),
		inmemory.NewContestRepository(seed.NewSeeds().Contests),
	)
	t.Run()
}

func TestFindProblemService_FindByID(t *testing.T) {
	act := []*problem.Data{
		func() *problem.Data {
			r, _ := pService.FindByID("110", time.Now())
			return r
		}(),
		func() *problem.Data {
			r, _ := pService.FindByID("110", time.Date(2020, time.April, 1, 12, 0, 0, 0, time.UTC))
			return r
		}(),
	}

	t.Run("コンテスト開始後は取得できる", func(t *testing.T) {
		assert.Equal(t, problem.DomainToData(seed.NewSeeds().Problems[0]), *act[0])
	})
	t.Run("コンテスト開始前は取得できない", func(t *testing.T) {
		assert.Equal(t, (*problem.Data)(nil), act[1])
	})

}

func TestFindProblemService_FindByContestID(t *testing.T) {

}
