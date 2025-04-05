package service

import (
	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/repository"
)

type SubmissionService struct {
	submissionRepository repository.SubmissionRepository
}

func NewSubmissionService(repository repository.SubmissionRepository) *SubmissionService {
	return &SubmissionService{submissionRepository: repository}
}

func (s *SubmissionService) IsExists(p domain.Submission) bool {
	// 重複判定: ID
	i, _ := s.submissionRepository.FindSubmissionByID(p.GetID())

	if i == nil {
		return false
	}

	return true
}
