package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
)

type SubmissionService struct {
	submissionRepository repository.SubmissionRepository
}

func NewSubmissionService(repository repository.SubmissionRepository) *SubmissionService {
	return &SubmissionService{submissionRepository: repository}
}

func (s *SubmissionService) IsExists(p domain.Submission) bool {
	// 重複判定: ID
	i := s.submissionRepository.FindSubmissionByID(p.GetID())

	if i == nil {
		return false
	}

	return true
}
