package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
)

type CaseSetService struct {
	problemRepository repository.ProblemRepository
}

func NewCaseSetService(repository repository.ProblemRepository) *CaseSetService {
	return &CaseSetService{problemRepository: repository}
}

func (s *CaseSetService) IsExists(p domain.Caseset) bool {
	// 重複判定: ID
	i, _ := s.problemRepository.FindCaseSetByID(p.GetID())

	if i == nil {
		return false
	}

	return true
}
