package service

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
)

type ContestService struct {
	contestRepository repository.ContestRepository
}

func NewContestService(repository repository.ContestRepository) *ContestService {
	return &ContestService{contestRepository: repository}
}

func (s *ContestService) IsExists(p model.Contest) bool {
	// 重複判定: ID/Title
	i, _ := s.contestRepository.FindContestByID(p.GetID())
	t, _ := s.contestRepository.FindContestByTitle(p.GetTitle())
	if i == nil && t == nil {
		return false
	}

	return true
}
