package contest

import (
	"errors"
	"fmt"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type JoinContestService struct {
	contestantRepository repository.ContestantRepository
	contestantService    service.ContestantService
}

func NewJoinContestService(contestantRepository repository.ContestantRepository, service service.ContestantService) *JoinContestService {
	return &JoinContestService{contestantRepository: contestantRepository, contestantService: service}
}

func (s JoinContestService) Join(contestID id.SnowFlakeID, user domain.User, role domain.ContestantRole) error {
	idGen := id.NewSnowFlakeIDGenerator()
	i := idGen.NewID(time.Now())

	d := domain.NewContestant(i, contestID, user.GetID())
	// システムの管理者は自動的にコンテストの管理者になる
	if user.IsAdmin() {
		d.SetAdmin()
	} else if role == domain.ContestTester {
		// テスターに指定されている場合はテスターとしてマークする
		d.SetTester()
	}

	if s.contestantService.IsExists(*d) {
		return errors.New("AlreadyJoined")
	}

	err := s.contestantRepository.JoinContest(*d)
	if err != nil {
		return fmt.Errorf("failed to join contest: %w", err)
	}
	return nil
}
