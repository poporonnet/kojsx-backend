package controller

import (
	"fmt"

	"github.com/poporonnet/kojsx-backend/pkg/contest/adaptor/controller/schema"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/contest/service/contest"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
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

func (c *ContestController) CreateContest(req schema.CreateContestRequestJSON) (schema.CreateContestResponseJSON, error) {
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
		return schema.CreateContestResponseJSON{}, fmt.Errorf("failed to create contest: %w", err)
	}
	return schema.CreateContestResponseJSON{
		ID:          string(res.GetID()),
		Title:       res.GetTitle(),
		Description: res.GetDescription(),
		StartAt:     res.GetStartAt(),
		EndAt:       res.GetEndAt(),
	}, nil
}

func (c *ContestController) FindContestByID(i string) (*schema.FindContestResponseJSON, error) {
	res, err := c.findService.FindByID(id.SnowFlakeID(i))
	if err != nil {
		return nil, fmt.Errorf("failed to find contest: %w", err)
	}
	return &schema.FindContestResponseJSON{
		ID:          string(res.GetID()),
		Title:       res.GetTitle(),
		Description: res.GetDescription(),
		StartAt:     res.GetStartAt(),
		EndAt:       res.GetEndAt(),
	}, nil
}

func (c *ContestController) FindContest() ([]schema.FindContestResponseJSON, error) {
	co, err := c.findService.FindAll()
	if err != nil {
		return nil, err
	}
	res := make([]schema.FindContestResponseJSON, len(co))
	for i, v := range co {
		res[i] = schema.FindContestResponseJSON{
			ID:          string(v.GetID()),
			Title:       v.GetTitle(),
			Description: v.GetDescription(),
			StartAt:     v.GetStartAt(),
			EndAt:       v.GetEndAt(),
		}
	}
	return res, nil
}

func (c *ContestController) GetRanking(i string) ([]schema.GetRankingResponseJSON, error) {
	res, err := c.getRankingService.Handle(id.SnowFlakeID(i))
	if err != nil {
		return nil, err
	}
	resp := make([]schema.GetRankingResponseJSON, len(res))
	for ii, v := range res {
		result := make([]schema.RankingProblemResult, len(v.Submissions))
		for j, k := range v.Submissions {
			result[j] = schema.RankingProblemResult{
				ProblemID: string(k.GetProblemID()),
				Point:     k.GetPoint(),
			}
		}
		resp[ii] = schema.GetRankingResponseJSON{
			Rank:  v.Rank,
			Point: v.Point,
			User: schema.RankingUser{
				ID:   string(v.User.GetID()),
				Name: v.User.GetName(),
			},
			Results: result,
		}
	}
	return resp, nil
}
