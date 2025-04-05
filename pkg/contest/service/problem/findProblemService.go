package problem

import (
	"errors"
	"fmt"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	repository2 "github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/utils"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type FindProblemService struct {
	repository repository2.ProblemRepository
	contest    repository2.ContestRepository
	contestant repository2.ContestantRepository
}

func NewFindProblemService(repo repository2.ProblemRepository, contest repository2.ContestRepository, contestantRepository repository2.ContestantRepository) *FindProblemService {
	return &FindProblemService{repo, contest, contestantRepository}
}

func (s *FindProblemService) FindByID(id id.SnowFlakeID, now time.Time, userID id.SnowFlakeID) (*Data, error) {
	p, err := s.repository.FindProblemByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find problem: %w", err)
	}
	co, err := s.contest.FindContestByID(p.GetContestID())
	if err != nil {
		return nil, fmt.Errorf("failed to find contest: %w", err)
	}
	ct, err := s.contestant.FindContestantByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to find contest: %w", err)
	}
	var role model.ContestantRole = model.ContestParticipants
	for _, v := range ct {
		if v.GetContestID() == co.GetID() {
			if v.IsAdmin() {
				role = model.ContestAdmin
			}
			if v.IsTester() {
				role = model.ContestTester
			}
			if v.IsNormal() {
				role = model.ContestParticipants
			}
			break
		}
	}
	// ToDo: リクエストしたユーザー(Contestant)の権限チェック
	if !co.IsStarted(now) && role == model.ContestParticipants {
		utils.SugarLogger.Errorf("contest id not started")
		return nil, errors.New("contest is not started")
	}
	res := DomainToData(*p)
	return &res, nil
}

func (s *FindProblemService) FindByContestID(id id.SnowFlakeID) ([]Data, error) {
	p, err := s.repository.FindProblemByContestID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find problem: %w", err)
	}
	co, err := s.contest.FindContestByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find contest: %w", err)
	}
	// ToDo: リクエストしたユーザー(Contestant)の権限チェック
	if !co.IsStarted(time.Now()) {
		utils.SugarLogger.Errorf("contest not started")
		return nil, errors.New("contest not started")
	}
	res := make([]Data, len(p))
	for i, v := range p {
		res[i] = DomainToData(v)
	}
	return res, nil
}
