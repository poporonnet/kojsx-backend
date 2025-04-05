package service

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
)

type CaseSetService struct {
	problemRepository repository.ProblemRepository
}

func NewCaseSetService(repository repository.ProblemRepository) *CaseSetService {
	return &CaseSetService{problemRepository: repository}
}

func (s *CaseSetService) IsExists(p model.Caseset) bool {
	// 重複判定: ID
	i, _ := s.problemRepository.FindCaseSetByID(p.GetID())

	return i != nil
}
