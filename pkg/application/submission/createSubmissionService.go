package submission

import (
	"errors"
	"fmt"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type CreateSubmissionService struct {
	repository  repository.SubmissionRepository
	service     service.SubmissionService
	idGenerator id.Generator
}

func NewCreateSubmissionService(
	repo repository.SubmissionRepository,
	service service.SubmissionService,
) *CreateSubmissionService {
	return &CreateSubmissionService{
		repo,
		service,
		id.NewSnowFlakeIDGenerator(),
	}
}

func (s CreateSubmissionService) Handle(pID, cID id.SnowFlakeID, lang, code string) (*domain.Submission, error) {
	newID := s.idGenerator.NewID(time.Now())

	d, err := domain.NewSubmission(newID, pID, cID, lang, code, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to create submission: %w", err)
	}

	if s.service.IsExists(*d) {
		return nil, errors.New("submission exists")
	}

	if err := s.repository.CreateSubmission(*d); err != nil {
		return nil, fmt.Errorf("failed to create submission: %w", err)
	}

	return d, nil
}
