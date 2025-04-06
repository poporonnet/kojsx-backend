package controller

import (
	"fmt"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/controller/schema"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/problem"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ProblemController struct {
	repository    repository.ProblemRepository
	createService problem.CreateProblemService
	findService   problem.FindProblemService
}

func NewProblemController(
	repository repository.ProblemRepository,
	createService problem.CreateProblemService,
	findService problem.FindProblemService,
) *ProblemController {
	return &ProblemController{
		repository:    repository,
		createService: createService,
		findService:   findService,
	}
}

func (c *ProblemController) CreateProblem(req schema.CreateProblemRequestJSON) (schema.CreateProblemResponseJSON, error) {
	res, err := c.createService.Handle(id.SnowFlakeID(req.ContestID), string(req.Title[0]), req.Title, req.Text, req.Points, req.Limits.Time)
	if err != nil {
		return schema.CreateProblemResponseJSON{}, fmt.Errorf("failed to create problem: %w", err)
	}
	return schema.CreateProblemResponseJSON{
		ID:     string(res.GetID()),
		Title:  res.GetTitle(),
		Text:   res.GetText(),
		Points: res.GetPoint(),
		Limits: struct {
			Memory int `json:"memory"`
			Time   int `json:"time"`
		}(struct {
			Memory int
			Time   int
		}{
			Memory: res.GetMemoryLimit(),
			Time:   res.GetTimeLimit(),
		}),
	}, nil
}

func (c *ProblemController) FindByID(i string) (schema.FindProblemResponseJSON, error) {
	// ToDo: 検索を実行したユーザーを渡す
	res, err := c.findService.FindByID(id.SnowFlakeID(i), time.Now(), "")
	if err != nil {
		return schema.FindProblemResponseJSON{}, fmt.Errorf("failed to find problem: %w", err)
	}
	return schema.CreateProblemResponseJSON{
		ID:     string(res.GetID()),
		Title:  res.GetTitle(),
		Text:   res.GetText(),
		Points: res.GetPoint(),
		Limits: struct {
			Memory int `json:"memory"`
			Time   int `json:"time"`
		}(struct {
			Memory int
			Time   int
		}{
			Memory: res.GetMemoryLimit(),
			Time:   res.GetTimeLimit(),
		}),
	}, nil
}

func (c *ProblemController) FindByContestID(id id.SnowFlakeID) ([]schema.FindProblemResponseJSON, error) {
	res, err := c.findService.FindByContestID(id)
	if err != nil {
		return []schema.FindProblemResponseJSON{}, fmt.Errorf("failed to find problem: %w", err)
	}

	response := make([]schema.FindProblemResponseJSON, len(res))
	for i, v := range res {
		response[i] = schema.CreateProblemResponseJSON{
			ID:     string(v.GetID()),
			Title:  v.GetTitle(),
			Text:   v.GetText(),
			Points: v.GetPoint(),
			Limits: struct {
				Memory int `json:"memory"`
				Time   int `json:"time"`
			}(struct {
				Memory int
				Time   int
			}{
				Memory: v.GetMemoryLimit(),
				Time:   v.GetTimeLimit(),
			}),
		}
	}

	return response, nil
}
