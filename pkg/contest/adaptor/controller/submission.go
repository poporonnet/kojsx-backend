package controller

import (
	"fmt"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/controller/schema"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/problem"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/submission"
	"github.com/poporonnet/kojsx-backend/pkg/user/service"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type SubmissionController struct {
	repository         repository.SubmissionRepository
	createService      submission.CreateSubmissionService
	findService        submission.FindSubmissionService
	findProblemService problem.FindProblemService
	findUserService    service.FindUserService
}

func NewSubmissionController(
	repository repository.SubmissionRepository,
	createService submission.CreateSubmissionService,
	findService submission.FindSubmissionService,
	findProblemService problem.FindProblemService,
	findUserService service.FindUserService,
) *SubmissionController {
	return &SubmissionController{
		repository,
		createService,
		findService,
		findProblemService,
		findUserService,
	}
}

func (c SubmissionController) CreateSubmission(cID string, req schema.CreateSubmissionRequestJSON) (schema.CreateSubmissionResponseJSON, error) {
	d, err := c.createService.Handle(
		id.SnowFlakeID(req.ProblemID),
		id.SnowFlakeID(cID),
		req.Lang,
		req.Code,
	)
	if err != nil {
		return schema.CreateSubmissionResponseJSON{}, err
	}

	return schema.CreateSubmissionResponseJSON{
		ID:        string(d.GetID()),
		ProblemID: string(d.GetProblemID()),
		Code:      d.GetCode(),
		Lang:      d.GetLang(),
	}, nil
}

func (c SubmissionController) FindByID(id id.SnowFlakeID) (schema.GetSubmissionResponseJSON, error) {
	s, err := c.findService.FindByID(id)
	if err != nil {
		return schema.GetSubmissionResponseJSON{}, err
	}
	// ToDo: 内部的な問題の取得のユーザーをどうするか
	p, err := c.findProblemService.FindByID(s.GetProblemID(), time.Now(), "")
	if err != nil {
		return schema.GetSubmissionResponseJSON{}, err
	}

	results := make([]schema.GetSubmissionResults, len(s.GetResults()))
	for i, v := range s.GetResults() {
		results[i] = schema.GetSubmissionResults{
			Name:   v.GetCaseName(),
			Status: v.GetResult(),
			Time:   v.GetExecTime(),
			Memory: v.GetExecMemory(),
		}
	}
	// ToDo: Contestant情報を入れる
	return schema.GetSubmissionResponseJSON{
		ID:          string(s.GetID()),
		SubmittedAt: s.GetSubmittedAt(),
		User: struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{},
		Problem: struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}{
			string(s.GetProblemID()),
			fmt.Sprintf("%s - %s", p.GetIndex(), p.GetTitle()),
		},
		Code:    s.GetCode(),
		Lang:    s.GetLang(),
		Points:  s.GetPoint(),
		Status:  s.GetResult(),
		Time:    s.GetExecTime(),
		Memory:  s.GetExecMemory(),
		Results: results,
	}, nil
}

func (c SubmissionController) CreateSubmissionResult(req schema.CreateSubmissionResultRequestJSON) error {
	arg := make([]submission.CreateResultArgs, len(req.Results))
	for i, v := range req.Results {
		arg[i] = submission.CreateResultArgs{
			Result:     "WJ",
			Output:     v.Output,
			CaseName:   v.CaseName,
			ExitStatus: v.ExitStatus,
			ExecTime:   v.Duration,
			ExecMemory: v.Usage,
		}
	}
	_, err := c.createService.CreateResult(id.SnowFlakeID(req.SubmissionID), arg)
	if err != nil {
		return err
	}
	return nil
}

func (c SubmissionController) FindTask() (schema.GetSubmissionTaskResponseJSON, error) {
	res, err := c.findService.FindTask()
	if err != nil {
		return schema.GetSubmissionTaskResponseJSON{}, err
	}

	// ToDo: Agent用のContestantを作る
	p, err := c.findProblemService.FindByID(res.GetProblemID(), time.Now(), "")
	if err != nil {
		return schema.GetSubmissionTaskResponseJSON{}, err
	}

	cases := make([]schema.GetSubmissionTaskResponseCases, 0)
	for _, v := range p.GetCaseSets() {
		for _, k := range v.GetCases() {
			cases = append(cases,
				schema.GetSubmissionTaskResponseCases{
					Name: fmt.Sprintf("%s.txt", k.GetID()),
					Data: k.GetIn(),
				},
			)
		}
	}
	return schema.GetSubmissionTaskResponseJSON{
		ID:        string(res.GetID()),
		ProblemID: string(res.GetProblemID()),
		Lang:      res.GetLang(),
		Code:      res.GetCode(),
		Cases:     cases,
		Config: schema.GetSubmissionTaskResponseConfig{
			TimeLimit:   p.GetTimeLimit(),
			MemoryLimit: p.GetMemoryLimit(),
		},
	}, nil
}

func (c SubmissionController) FindByContestID(i string) ([]schema.FindSubmissionByContestIDResponseJSON, error) {
	r, err := c.findService.FindByContestID(id.SnowFlakeID(i))
	if err != nil {
		return nil, err
	}

	res := make([]schema.FindSubmissionByContestIDResponseJSON, 0)
	for _, k := range r.S {
		problemName := ""
		for _, q := range r.P {
			if k.GetProblemID() == q.GetID() {
				problemName = fmt.Sprintf("%s - %s", q.GetIndex(), q.GetTitle())
			}
		}
		res = append(res, schema.FindSubmissionByContestIDResponseJSON{
			ID:          string(k.GetID()),
			SubmittedAt: k.GetSubmittedAt(),
			User: struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			}{
				string(k.GetContestantID()),
				"",
			},
			Problem: struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			}{
				string(k.GetProblemID()),
				problemName,
			},
			Lang:   k.GetLang(),
			Points: k.GetPoint(),
			Status: k.GetResult(),
			Time:   k.GetExecTime(),
			Memory: k.GetExecMemory(),
		},
		)
	}
	return res, nil
}
