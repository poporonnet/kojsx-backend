package problem

import (
	"errors"
	"fmt"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type FindProblemService struct {
	repository repository.ProblemRepository
	contest    repository.ContestRepository
}

func NewFindProblemService(repo repository.ProblemRepository, contest repository.ContestRepository) *FindProblemService {
	return &FindProblemService{repo, contest}
}

func (s *FindProblemService) FindByID(id id.SnowFlakeID, now time.Time) (*Data, error) {
	p, err := s.repository.FindProblemByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find problem: %w", err)
	}
	co, err := s.contest.FindContestByID(p.GetContestID())
	if err != nil {
		return nil, fmt.Errorf("failed to find contest: %w", err)
	}
	// ToDo: リクエストしたユーザー(Contestant)の権限チェック
	if !co.IsStarted(now) {
		utils.SugarLogger.Errorf("contest id not started")
		return nil, errors.New("contest is not started")
	}
	res := DomainToData(*p)
	return &res, nil
}

func (s *FindProblemService) FindByContestID(id id.SnowFlakeID) ([]Data, error) {
	p, err := s.repository.FindProblemByContestID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find problem: %w", err)
	}
	co, err := s.contest.FindContestByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find contest: %w", err)
	}
	// ToDo: リクエストしたユーザー(Contestant)の権限チェック
	if !co.IsStarted(time.Now()) {
		utils.SugarLogger.Errorf("contest not started")
		return nil, errors.New("contest not started")
	}
	res := make([]Data, len(p))
	for i, v := range p {
		res[i] = DomainToData(v)
	}
	return res, nil
}
