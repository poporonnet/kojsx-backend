package submission_test

import (
	"testing"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/repository/inmemory"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/service"
	submission2 "github.com/poporonnet/kojsx-backend/pkg/contest/service/submission"

	"github.com/poporonnet/kojsx-backend/pkg/utils"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"
	"github.com/stretchr/testify/assert"
)

var ss *submission2.CreateSubmissionService
var loc = time.UTC
var d = time.Date(2021, time.October, 1, 0, 0, 0, 0, loc)

func TestMain(m *testing.M) {
	utils.NewLogger()
	s := seed.NewSeeds()
	submissionRepository := inmemory.NewSubmissionRepository(s.Submission)
	problemRepository := inmemory.NewProblemRepository(s.Problems)
	ss = submission2.NewCreateSubmissionService(submissionRepository, *service.NewSubmissionService(submissionRepository), problemRepository)
	m.Run()
}

func TestCreateSubmissionService_CreateResult(t *testing.T) {
	args := []submission2.CreateResultArgs{
		{
			Result:     "WJ",
			Output:     "world\n",
			CaseName:   "70",
			ExitStatus: 0,
			ExecTime:   10,
			ExecMemory: 900,
		},
		{
			Result:     "WJ",
			Output:     "3\n",
			CaseName:   "80",
			ExitStatus: 0,
			ExecTime:   20,
			ExecMemory: 850,
		},
		{
			Result:     "WJ",
			Output:     "3\n",
			CaseName:   "90",
			ExitStatus: 0,
			ExecTime:   30,
			ExecMemory: 900,
		},
		{
			Result:     "WJ",
			Output:     "2\n",
			CaseName:   "100",
			ExitStatus: 0,
			ExecTime:   40,
			ExecMemory: 900,
		},
	}
	res, _ := ss.CreateResult("200", args)
	e, _ := model.NewSubmission(
		"200",
		"110",
		"1",
		"G++",
		"I2luY2x1ZGUgPGlvc3RyZWFtPgoKaW50IG1haW4oKSB7CiAgc3RkOjpjb3V0IDw8ICJ3b3JsZCIgPDwgc3RkOjplbmRsOwp9Cg==",
		d,
	)
	e.SetResult("AC")
	_ = e.SetPoint(300)
	e.SetExecTime(40)
	e.SetExecMemory(900)

	t.Run("Result以外のテスト", func(t *testing.T) {
		// Result以外のテスト
		act := submission2.NewData(res.GetID(), res.GetProblemID(), res.GetContestantID(), res.GetPoint(), res.GetLang(), res.GetCodeLength(), res.GetResult(), res.GetExecTime(), res.GetExecMemory(), res.GetCode(), res.GetSubmittedAt(), []submission2.Result{})
		assert.Equal(t, submission2.DomainToData(*e), act)
	})

	t.Run("Resultのテスト", func(t *testing.T) {
		act := res.GetResults()
		for i, tt := range args {
			t.Run(tt.CaseName, func(t *testing.T) {
				exp := *submission2.NewResult(
					act[i].GetID(),
					tt.Output,
					"AC",
					tt.CaseName,
					tt.ExitStatus,
					tt.ExecTime,
					tt.ExecMemory,
				)
				assert.Equal(t, exp, act[i])
			})
		}
	})
}
