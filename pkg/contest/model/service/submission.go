package service

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
)

type SubmissionService struct {
	submissionRepository repository.SubmissionRepository
}

func NewSubmissionService(repository repository.SubmissionRepository) *SubmissionService {
	return &SubmissionService{submissionRepository: repository}
}

func (s *SubmissionService) IsExists(p model.Submission) bool {
	// 重複判定: ID
	i, _ := s.submissionRepository.FindSubmissionByID(p.GetID())

	return i != nil
}
