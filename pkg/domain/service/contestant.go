package service

import (
	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/repository"
)

type ContestantService struct {
	contestantRepository repository.ContestantRepository
}

func NewContestantService(repository repository.ContestantRepository) *ContestantService {
	return &ContestantService{contestantRepository: repository}
}

func (s *ContestantService) IsExists(p domain.Contestant) bool {
	// 重複判定: ID/UserID/ContestID
	i, _ := s.contestantRepository.FindContestantByID(p.GetID())
	u, _ := s.contestantRepository.FindContestantByUserID(p.GetUserID())
	c, _ := s.contestantRepository.FindContestantByContestID(p.GetContestID())
	/*
		IDはいかなる場合でも重複してはいけない
		UserIDとContestIDの組み合わせが存在するときは重複しているとみなす
	*/
	// IDが重複していない
	if i == nil {
		return false
	}

	//
	for _, v := range c {
		for _, k := range u {
			if v.GetContestID() == p.GetContestID() && k.GetUserID() == p.GetUserID() {
				return true
			}
		}
	}

	return false
}
