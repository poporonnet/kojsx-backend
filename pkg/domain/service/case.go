package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
)

type CaseService struct {
	problemRepository repository.ProblemRepository
}

func NewCaseService(repository repository.ProblemRepository) *CaseService {
	return &CaseService{problemRepository: repository}
}

func (s *CaseService) IsExists(p domain.Case) bool {
	// 重複判定: ID
	i := s.problemRepository.FindCaseByID(p.GetID())

	if i == nil {
		return false
	}

	return true
}
