package problem

import (
	"errors"
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"time"
)

type CreateProblemService struct {
	problemRepository repository.ProblemRepository
	problemService    service.ProblemService
}

func NewCreateProblemService(repository repository.ProblemRepository, service service.ProblemService) *CreateProblemService {
	return &CreateProblemService{
		problemService:    service,
		problemRepository: repository,
	}
}

func (s *CreateProblemService) Handle(
	contestID id.SnowFlakeID,
	index, title, text string,
	point, timeLimit int,
) (*Data, error) {
	gen := id.NewSnowFlakeIDGenerator()
	id := gen.NewID(time.Now())
	p := domain.NewProblem(id, contestID)
	err := p.SetIndex(index)
	if err != nil {
		return nil, err
	}
	err = p.SetTitle(title)
	if err != nil {
		return nil, err
	}
	err = p.SetText(text)
	if err != nil {
		return nil, err
	}
	err = p.SetPoint(point)
	if err != nil {
		return nil, err
	}
	err = p.SetTimeLimit(timeLimit)
	if err != nil {
		return nil, err
	}

	if s.problemService.IsExists(*p) {
		return nil, errors.New("already exists")
	}

	if err := s.problemRepository.CreateProblem(*p); err != nil {
		return nil, err
	}

	res := DomainToData(*p)

	return &res, nil
}
