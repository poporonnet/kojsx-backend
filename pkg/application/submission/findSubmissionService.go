package submission

import (
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type FindSubmissionService struct {
	submissionRepository repository.SubmissionRepository
}

func NewFindSubmissionService(submissionRepository repository.SubmissionRepository) *FindSubmissionService {
	return &FindSubmissionService{submissionRepository: submissionRepository}
}

func (s FindSubmissionService) FindByID(id id.SnowFlakeID) (*Data, error) {
	su, _ := s.submissionRepository.FindSubmissionByID(id)
	return DomainToData(*su), nil
}
