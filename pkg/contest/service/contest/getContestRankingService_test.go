package contest_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/contest"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/problem"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/submission"
	userInmemory "github.com/poporonnet/kojsx-backend/pkg/user/adaptor/repository/inmemory"
	userModel "github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
	"github.com/stretchr/testify/assert"
)

var s contest.GetContestRankingService

func TestMain(t *testing.M) {
	contests := []model.Contest{
		contest.NewData(
			"1",
			"Test Contest",
			"Hello world!",
			time.Now().Add(10),
			time.Now().Add(1*time.Hour),
		).ToDomain(),
	}

	sets := []problem.CaseSetData{
		*problem.NewCaseSetData(
			"100",
			"001",
			100,
			[]problem.CaseData{
				*problem.NewCaseData(
					"10",
					"100",
					"hello",
					"world",
				),
				*problem.NewCaseData(
					"11",
					"100",
					"moyo",
					"boge",
				),
			},
		),
		*problem.NewCaseSetData(
			"101",
			"002",
			100,
			[]problem.CaseData{
				*problem.NewCaseData(
					"12",
					"101",
					"hoge",
					"fuga",
				),
				*problem.NewCaseData(
					"13",
					"101",
					"piyo",
					"gebo",
				),
			},
		),
	}

	problems := []model.Problem{
		func() model.Problem {
			p := problem.NewData(
				"1000",
				"1",
				"A",
				"hello world",
				"print hello world",
				100,
				300,
				sets,
			)
			return *p.ToDomain()
		}(),
	}

	users := []userModel.User{
		service.NewData(
			"10000",
			"Eric",
			"eric@example.com",
			"",
			userModel.Admin,
		).ToDomain(),
		service.NewData(
			"10001",
			"George",
			"george@example.com",
			"",
			userModel.Normal,
		).ToDomain(),
		service.NewData(
			"10002",
			"Joan",
			"joan@example.com",
			"",
			userModel.Normal,
		).ToDomain(),
		service.NewData(
			"10003",
			"Kate",
			"kate@example.com",
			"",
			userModel.Normal,
		).ToDomain(),
	}

	contestants := []model.Contestant{
		*model.NewContestant(
			"100000",
			"1",
			"10000",
		),
		*model.NewContestant(
			"100001",
			"1",
			"10001",
		),
		*model.NewContestant(
			"100002",
			"1",
			"10002",
		),
		*model.NewContestant(
			"100003",
			"1",
			"10003",
		),
	}
	contestants[0].SetAdmin()
	contestants[1].SetTester()
	contestants[2].SetNormal()
	contestants[3].SetNormal()

	submissions := []model.Submission{
		*submission.NewData(
			"0",
			"1000",
			"100000",
			200,
			"Ruby",
			10,
			"AC",
			100,
			1000,
			"p `echo hello world`",
			time.Now().Add(300),
			[]submission.Result{
				*submission.NewResult(
					"00",
					"world",
					"AC",
					"10",
					0,
					100,
					100,
				),
			},
		).ToDomain(),
		*submission.NewData(
			"01",
			"1000",
			"100001",
			150,
			"Ruby",
			10,
			"WA",
			100,
			1000,
			"p `echo hello world`",
			time.Now().Add(300),
			[]submission.Result{
				*submission.NewResult(
					"00",
					"world",
					"AC",
					"10",
					0,
					100,
					100,
				),
			},
		).ToDomain(),
		*submission.NewData(
			"02",
			"1000",
			"100002",
			140,
			"Ruby",
			10,
			"TLE",
			100,
			1000,
			"p `echo hello world`",
			time.Now().Add(300),
			[]submission.Result{
				*submission.NewResult(
					"00",
					"world",
					"AC",
					"10",
					0,
					1000000000,
					100,
				),
			},
		).ToDomain(),
		*submission.NewData(
			"03",
			"1000",
			"100003",
			130,
			"Ruby",
			10,
			"AC",
			100,
			1000,
			"p `echo hello world`",
			time.Now().Add(300),
			[]submission.Result{
				*submission.NewResult(
					"00",
					"world",
					"AC",
					"10",
					0,
					100,
					100,
				),
			},
		).ToDomain(),
		*submission.NewData(
			"04",
			"1000",
			"100003",
			0,
			"Ruby",
			10,
			"WA",
			100,
			1000,
			"p `echo hello world`",
			time.Now().Add(300),
			[]submission.Result{
				*submission.NewResult(
					"00",
					"world",
					"AC",
					"10",
					0,
					100,
					100,
				),
			},
		).ToDomain(),
	}
	fmt.Println(submissions)

	s = *contest.NewGetContestRankingService(
		inmemory.NewContestRepository(contests),
		inmemory.NewContestantRepository(contestants),
		inmemory.NewProblemRepository(problems),
		inmemory.NewSubmissionRepository(submissions),
		userInmemory.NewUserRepository(users),
	)
	utils.NewLogger()
	t.Run()
}

func TestGetContestRankingService_Handle(t *testing.T) {
	type exp struct {
		Rank         int
		Point        int
		ContestantID id.SnowFlakeID
	}
	expect := []exp{
		{
			1,
			140,
			"100002",
		},
		{
			2,
			130,
			"100003",
		},
	}
	/*
		想定結果
		0. Contestant No. 100000 200pts (Admin)
		0. Contestant No. 100001 150pts (Tester)
		1. Contestant No. 100002 140pts
		2. Contestant No. 100003 130pts (1 attempt)
	*/
	t.Run("正常に順位の計算ができる", func(t *testing.T) {
		res, _ := s.Handle("1")

		for i, tt := range res {
			t.Run(fmt.Sprintf("%d位が正しいか", i+1), func(t *testing.T) {
				act := exp{
					tt.Rank,
					tt.Point,
					tt.Contestant.GetID(),
				}
				assert.Equal(t, expect[i], act)
			})
		}
	})
}
