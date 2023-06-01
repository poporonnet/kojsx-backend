package problem

import (
	"errors"
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
	p := s.repository.FindProblemByID(id)
	if p == nil {
		return nil, errors.New("not found")
	}
	res := DomainToData(*p)
	return &res, nil
}

func (s *FindProblemService) FindByContestID(id id.SnowFlakeID) ([]Data, error) {
	p := s.repository.FindProblemByContestID(id)
	res := make([]Data, len(p))
	for i, v := range p {
		res[i] = DomainToData(v)
	}
	return res, nil
}
