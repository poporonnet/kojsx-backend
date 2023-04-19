package service

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
)

type ContestService struct {
	contestRepository repository.ContestRepository
}

func NewContestService(repository repository.ContestRepository) *ContestService {
	return &ContestService{contestRepository: repository}
}

func (s *ContestService) IsExists(p domain.Contest) bool {
	// 重複判定: ID/Title
	i := s.contestRepository.FindContestByID(p.GetID())
	t := s.contestRepository.FindContestByTitle(p.GetTitle())
	if i == nil && t == nil {
		return false
	}

	return true
}
