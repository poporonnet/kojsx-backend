package problem

import (
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type FindProblemService struct {
	repository repository.ProblemRepository
}

func NewFindProblemService(repo repository.ProblemRepository) *FindProblemService {
	return &FindProblemService{repo}
}

func (s *FindProblemService) FindByID(id id.SnowFlakeID) (*Data, error) {
	p, err := s.repository.FindProblemByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find problem: %w", err)
	}
	res := DomainToData(*p)
	return &res, nil
}

func (s *FindProblemService) FindByContestID(id id.SnowFlakeID) ([]Data, error) {
	p, err := s.repository.FindProblemByContestID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find problem: %w", err)
	}
	res := make([]Data, len(p))
	for i, v := range p {
		res[i] = DomainToData(v)
	}
	return res, nil
}
