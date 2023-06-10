package controller

import (
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/application/problem"
	"github.com/mct-joken/kojs5-backend/pkg/application/submission"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type SubmissionController struct {
	repository         repository.SubmissionRepository
	createService      submission.CreateSubmissionService
	findService        submission.FindSubmissionService
	findProblemService problem.FindProblemService
}

func NewSubmissionController(
	repository repository.SubmissionRepository,
	createService submission.CreateSubmissionService,
	findService submission.FindSubmissionService,
	findProblemService problem.FindProblemService,
) *SubmissionController {
	return &SubmissionController{
		repository,
		createService,
		findService,
		findProblemService,
	}
}

func (c SubmissionController) CreateSubmission(cID string, req model.CreateSubmissionRequestJSON) (model.CreateSubmissionResponseJSON, error) {
	d, err := c.createService.Handle(
		id.SnowFlakeID(req.ProblemID),
		id.SnowFlakeID(cID),
		req.Lang,
		req.Code,
	)
	if err != nil {
		return model.CreateSubmissionResponseJSON{}, err
	}

	return model.CreateSubmissionResponseJSON{
		ID:        string(d.GetID()),
		ProblemID: string(d.GetProblemID()),
		Code:      d.GetCode(),
		Lang:      d.GetLang(),
	}, nil
}

func (c SubmissionController) CreateSubmissionResult(req model.CreateSubmissionResultRequestJSON) error {
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
	err := c.createService.CreateResult(id.SnowFlakeID(req.SubmissionID), arg)
	if err != nil {
		return err
	}
	return nil
}

func (c SubmissionController) FindTask() (model.GetSubmissionTaskResponseJSON, error) {
	res, err := c.findService.FindTask()
	if err != nil {
		return model.GetSubmissionTaskResponseJSON{}, err
	}

	p, err := c.findProblemService.FindByID(res.GetProblemID())
	if err != nil {
		return model.GetSubmissionTaskResponseJSON{}, err
	}

	cases := make([]model.GetSubmissionTaskResponseCases, 0)
	for _, v := range p.GetCaseSets() {
		for _, k := range v.GetCases() {
			cases = append(cases,
				model.GetSubmissionTaskResponseCases{
					Name: fmt.Sprintf("%s.txt", v.GetID()),
					Data: k.GetIn(),
				},
			)
		}
	}

	return model.GetSubmissionTaskResponseJSON{
		ID:        string(res.GetID()),
		ProblemID: string(res.GetProblemID()),
		Lang:      res.GetLang(),
		Code:      res.GetCode(),
		Cases:     cases,
		Config: model.GetSubmissionTaskResponseConfig{
			TimeLimit:   p.GetTimeLimit(),
			MemoryLimit: p.GetMemoryLimit(),
		},
	}, nil
}
