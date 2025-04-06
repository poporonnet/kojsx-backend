package seed

import (
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/contest"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/problem"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/submission"
	model2 "github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/user/service"
)

type Seeds struct {
	Contests    []model.Contest
	Contestants []model.Contestant
	Users       []model2.User
	Problems    []model.Problem
	Submission  []model.Submission
}

func NewSeeds() Seeds {
	loc := time.UTC
	d := time.Date(2021, time.October, 1, 0, 0, 0, 0, loc)

	contests := func() []model.Contest {
		data := contest.NewData(
			"10",
			"Contest 1",
			"# About\nThis Contest is for seed",
			d,
			d.Add(24*time.Hour*31*12*10),
		)
		return []model.Contest{data.ToDomain()}
	}()

	contestants := func() []model.Contestant {
		data := model.NewContestant("900", "10", "20")
		data.SetAdmin()
		data2 := model.NewContestant("910", "10", "30")
		data2.SetNormal()
		return []model.Contestant{*data, *data2}
	}()

	users := func() []model2.User {
		user1 := service.NewData(
			"20",
			"Eric",
			"eric@example.jp",
			"Argon2.8ce04ed8562b03c813343a04022f93db7629f9f2.1a7[0]",
			model2.Admin,
		)

		user2 := service.NewData(
			"30",
			"Eric",
			"eric@example.jp",
			"Argon2.8ce04ed8562b03c813343a04022f93db7629f9f2.1a7[0]",
			model2.Normal,
		)
		return []model2.User{user1.ToDomain(), user2.ToDomain()}
	}()

	problems := func() []model.Problem {
		// Case
		case1 := *problem.NewCaseData(
			"70",
			"50",
			"hello\n",
			"world\n",
		)
		case2 := *problem.NewCaseData(
			"80",
			"50",
			"1 2\n",
			"3\n",
		)
		case3 := *problem.NewCaseData(
			"90", "60",
			"abc\n",
			"3\n",
		)
		case4 := *problem.NewCaseData(
			"100", "60",
			"abc\nabp abc\n",
			"2\n",
		)

		// CaseSets
		set1 := *problem.NewCaseSetData(
			"50",
			"test1",
			200,
			[]problem.CaseData{
				case1, case2,
			},
		)
		set2 := *problem.NewCaseSetData(
			"60",
			"test2",
			100,
			[]problem.CaseData{
				case3, case4,
			},
		)

		// Problems
		problem1 := problem.NewData(
			"110",
			"10",
			"A",
			"Moji",
			"Calculate the number.\n",
			300,
			2000,
			[]problem.CaseSetData{
				set1, set2,
			},
		)

		return []model.Problem{*problem1.ToDomain()}
	}()

	submissions := func() []model.Submission {
		submission1 := submission.NewData(
			"200",
			problems[0].GetProblemID(),
			"1",
			0,
			"G++",
			180,
			"WE",
			0,
			0,
			"I2luY2x1ZGUgPGlvc3RyZWFtPgoKaW50IG1haW4oKSB7CiAgc3RkOjpjb3V0IDw8ICJ3b3JsZCIgPDwgc3RkOjplbmRsOwp9Cg==",
			d,
			nil,
		)
		submission2 := submission.NewData(
			"210",
			problems[0].GetProblemID(),
			"1",
			0,
			"G++",
			180,
			"WE",
			0,
			0,
			"I2luY2x1ZGUgPGlvc3RyZWFtPgoKaW50IG1haW4oKSB7CiAgc3RkOjpjb3V0IDw8ICJ3b3JsZCIgPDwgc3RkOjplbmRsOwp9Cg==",
			d,
			nil,
		)
		return []model.Submission{*submission1.ToDomain(), *submission2.ToDomain()}
	}()

	return Seeds{
		Contests:    contests,
		Contestants: contestants,
		Users:       users,
		Problems:    problems,
		Submission:  submissions,
	}
}
