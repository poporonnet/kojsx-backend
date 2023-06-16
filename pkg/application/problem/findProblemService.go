package problem

import (
	"errors"
	"fmt"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"

	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type FindProblemService struct {
	repository repository.ProblemRepository
	contest    repository.ContestRepository
	contestant repository.ContestantRepository
}

func NewFindProblemService(repo repository.ProblemRepository, contest repository.ContestRepository, contestantRepository repository.ContestantRepository) *FindProblemService {
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
	var role domain.ContestantRole = domain.ContestParticipants
	for _, v := range ct {
		if v.GetContestID() == co.GetID() {
			if v.IsAdmin() {
				role = domain.ContestAdmin
			}
			if v.IsTester() {
				role = domain.ContestTester
			}
			if v.IsNormal() {
				role = domain.ContestParticipants
			}
			break
		}
	}
	// ToDo: リクエストしたユーザー(Contestant)の権限チェック
	if !co.IsStarted(now) && role == domain.ContestParticipants {
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
