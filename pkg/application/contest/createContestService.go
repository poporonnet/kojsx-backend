package contest

import (
	"errors"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type CreateContestService struct {
	contestRepository repository.ContestRepository
	contestService    *service.ContestService
}

func NewCreateContestService(contestRepository repository.ContestRepository) *CreateContestService {
	return &CreateContestService{
		contestRepository: contestRepository,
		contestService:    service.NewContestService(contestRepository),
	}
}

func (s *CreateContestService) Handle(title string, description string, startAt time.Time, endAt time.Time) (*Data, error) {
	gen := id.NewSnowFlakeIDGenerator()
	id := gen.NewID(time.Now())
	c := domain.NewContest(id)
	if err := c.SetTitle(title); err != nil {
		return nil, err
	}
	if err := c.SetDescription(description); err != nil {
		return nil, err
	}
	if err := c.SetStartAt(startAt); err != nil {
		return nil, err
	}
	if err := c.SetEndAt(endAt); err != nil {
		return nil, err
	}

	if s.contestService.IsExists(*c) {
		return nil, errors.New("AlreadyExists")
	}

	if err := s.contestRepository.CreateContest(*c); err != nil {
		return nil, err
	}
	r := DomainToData(*c)
	return &r, nil
}
