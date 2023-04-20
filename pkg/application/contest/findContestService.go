package contest

import (
	"errors"

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
	r := s.contestRepository.FindContestByID(id)
	if r == nil {
		return nil, errors.New("NotExists")
	}
	res := DomainToData(*r)
	return &res, nil
}

func (s *FindContestService) FindAll() ([]Data, error) {
	r := s.contestRepository.FindAllContests()
	res := make([]Data, len(r))
	for i, v := range r {
		res[i] = DomainToData(v)
	}
	return res, nil
}
