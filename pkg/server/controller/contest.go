package controller

import (
	"github.com/mct-joken/kojs5-backend/pkg/application/contest"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
)

type ContestController struct {
	repository    repository.ContestRepository
	createService contest.CreateContestService
}

func NewContestController(repository repository.ContestRepository, createService contest.CreateContestService) *ContestController {
	return &ContestController{
		repository:    repository,
		createService: createService,
	}
}

func (c *ContestController) CreateContest(req model.CreateContestRequestJSON) (model.CreateContestResponseJSON, error) {
	res, err := c.createService.Handle(req.Title, req.Description, req.StartAt, req.EndAt)
	if err != nil {
		return model.CreateContestResponseJSON{}, err
	}
	return model.CreateContestResponseJSON{
		ID:          string(res.GetID()),
		Title:       res.GetTitle(),
		Description: res.GetDescription(),
		StartAt:     res.GetStartAt(),
		EndAt:       res.GetEndAt(),
	}, nil
}
