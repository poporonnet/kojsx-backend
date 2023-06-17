package controller

import (
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/application/contest"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ContestController struct {
	repository        repository.ContestRepository
	createService     contest.CreateContestService
	findService       contest.FindContestService
	getRankingService contest.GetContestRankingService
}

func NewContestController(
	repository repository.ContestRepository,
	createService contest.CreateContestService,
	findService contest.FindContestService,
	service contest.GetContestRankingService,
) *ContestController {
	return &ContestController{
		repository:        repository,
		createService:     createService,
		findService:       findService,
		getRankingService: service,
	}
}

func (c *ContestController) CreateContest(req model.CreateContestRequestJSON) (model.CreateContestResponseJSON, error) {
	res, err := c.createService.Handle(
		contest.CreateContestArgs{
			Title:       req.Title,
			Description: req.Description,
			StartAt:     req.StartAt,
			EndAt:       req.EndAt,
			// ToDo: 操作を行うユーザーを指定する
		},
	)
	if err != nil {
		return model.CreateContestResponseJSON{}, fmt.Errorf("failed to create contest: %w", err)
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
		return nil, fmt.Errorf("failed to find contest: %w", err)
	}
	return &model.FindContestResponseJSON{
		ID:          string(res.GetID()),
		Title:       res.GetTitle(),
		Description: res.GetDescription(),
		StartAt:     res.GetStartAt(),
		EndAt:       res.GetEndAt(),
	}, nil
}

func (c *ContestController) FindContest() ([]model.FindContestResponseJSON, error) {
	co, err := c.findService.FindAll()
	if err != nil {
		return nil, err
	}
	res := make([]model.FindContestResponseJSON, len(co))
	for i, v := range co {
		res[i] = model.FindContestResponseJSON{
			ID:          string(v.GetID()),
			Title:       v.GetTitle(),
			Description: v.GetDescription(),
			StartAt:     v.GetStartAt(),
			EndAt:       v.GetEndAt(),
		}
	}
	return res, nil
}

func (c *ContestController) GetRanking(i string) ([]model.GetRankingResponseJSON, error) {
	res, err := c.getRankingService.Handle(id.SnowFlakeID(i))
	if err != nil {
		return nil, err
	}
	resp := make([]model.GetRankingResponseJSON, len(res))
	for ii, v := range res {
		result := make([]model.RankingProblemResult, len(v.Submissions))
		for j, k := range v.Submissions {
			result[j] = model.RankingProblemResult{
				ProblemID: string(k.GetProblemID()),
				Point:     k.GetPoint(),
			}
		}
		resp[ii] = model.GetRankingResponseJSON{
			Rank:  v.Rank,
			Point: v.Point,
			User: model.RankingUser{
				ID:   string(v.User.GetID()),
				Name: v.User.GetName(),
			},
			Results: result,
		}
	}
	return resp, nil
}
