package service

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
)

type CaseService struct {
	problemRepository repository.ProblemRepository
}

func NewCaseService(repository repository.ProblemRepository) *CaseService {
	return &CaseService{problemRepository: repository}
}

func (s *CaseService) IsExists(p model.Case) bool {
	// 重複判定: ID
	i, _ := s.problemRepository.FindCaseByID(p.GetID())

	return i != nil
}
