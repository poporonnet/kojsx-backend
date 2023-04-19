package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
)

type SubmissionResultService struct {
	submissionRepository repository.SubmissionRepository
}

func NewSubmissionResultService(repository repository.SubmissionRepository) *SubmissionResultService {
	return &SubmissionResultService{submissionRepository: repository}
}

func (s *SubmissionResultService) IsExists(p domain.SubmissionResult) bool {
	// 重複判定: ID
	i := s.submissionRepository.FindSubmissionResultByID(p.GetID())

	if i == nil {
		return false
	}

	return true
}
