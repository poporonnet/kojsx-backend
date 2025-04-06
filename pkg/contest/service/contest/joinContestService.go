package contest

import (
	"errors"
	"fmt"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/repository"
	"github.com/poporonnet/kojsx-backend/pkg/contest/model/service"
	userModel "github.com/poporonnet/kojsx-backend/pkg/user/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type JoinContestService struct {
	contestantRepository repository.ContestantRepository
	contestantService    service.ContestantService
}

func NewJoinContestService(contestantRepository repository.ContestantRepository, service service.ContestantService) *JoinContestService {
	return &JoinContestService{contestantRepository: contestantRepository, contestantService: service}
}

func (s JoinContestService) Join(contestID id.SnowFlakeID, user userModel.User, role model.ContestantRole) error {
	idGen := id.NewSnowFlakeIDGenerator()
	i := idGen.NewID(time.Now())

	d := model.NewContestant(i, contestID, user.GetID())
	// システムの管理者は自動的にコンテストの管理者になる
	if user.IsAdmin() {
		d.SetAdmin()
	} else if role == model.ContestTester {
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
