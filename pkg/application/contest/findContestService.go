package contest

import (
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type FindContestService struct {
	contestRepository repository.ContestRepository
}

func NewFindContestService(contestRepository repository.ContestRepository) *FindContestService {
	return &FindContestService{contestRepository: contestRepository}
}

func (s *FindContestService) FindByID(id id.SnowFlakeID) (*Data, error) {
	r, err := s.contestRepository.FindContestByID(id)
	if err != nil {
		return nil, fmt.Errorf("not found: %w", err)
	}
	res := DomainToData(*r)
	return &res, nil
}

func (s *FindContestService) FindAll() ([]Data, error) {
	r, err := s.contestRepository.FindAllContests()
	if err != nil {
		return nil, fmt.Errorf("failed to find all users: %w", err)
	}
	res := make([]Data, len(r))
	for i, v := range r {
		res[i] = DomainToData(v)
	}
	return res, nil
}
