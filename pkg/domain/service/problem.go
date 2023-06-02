package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
)

type ProblemService struct {
	problemRepository repository.ProblemRepository
}

func NewProblemService(repository repository.ProblemRepository) *ProblemService {
	return &ProblemService{problemRepository: repository}
}

func (s *ProblemService) IsExists(p domain.Problem) bool {
	// 重複判定: ID/title
	i, _ := s.problemRepository.FindProblemByID(p.GetProblemID())
	t, _ := s.problemRepository.FindProblemByTitle(p.GetTitle())

	if i == nil && t == nil {
		return false
	}

	return true
}
