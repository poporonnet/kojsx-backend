package service

import (
	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/repository"
)

type CaseService struct {
	problemRepository repository.ProblemRepository
}

func NewCaseService(repository repository.ProblemRepository) *CaseService {
	return &CaseService{problemRepository: repository}
}

func (s *CaseService) IsExists(p domain.Case) bool {
	// 重複判定: ID
	i, _ := s.problemRepository.FindCaseByID(p.GetID())

	if i == nil {
		return false
	}

	return true
}
