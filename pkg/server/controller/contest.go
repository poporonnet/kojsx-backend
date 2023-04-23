package controller

import (
	"github.com/mct-joken/kojs5-backend/pkg/application/contest"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ContestController struct {
	repository    repository.ContestRepository
	createService contest.CreateContestService
	findService   contest.FindContestService
}

func NewContestController(repository repository.ContestRepository, createService contest.CreateContestService, findService contest.FindContestService) *ContestController {
	return &ContestController{
		repository:    repository,
		createService: createService,
		findService:   findService,
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

func (c *ContestController) FindContestByID(i string) (*model.FindContestResponseJSON, error) {
	res, err := c.findService.FindByID(id.SnowFlakeID(i))
	if err != nil {
		return nil, err
	}
	return &model.FindContestResponseJSON{
		ID:          string(res.GetID()),
		Title:       res.GetTitle(),
		Description: res.GetDescription(),
		StartAt:     res.GetStartAt(),
		EndAt:       res.GetEndAt(),
	}, nil
}
