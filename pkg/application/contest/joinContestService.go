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

func (s JoinContestService) Join(contestID, userID id.SnowFlakeID) error {
	i := id.NewSnowFlakeIDGenerator()
	id := i.NewID(time.Now())

	d := domain.NewContestant(id, contestID, userID)

	if s.contestantService.IsExists(*d) {
		return errors.New("AlreadyJoined")
	}

	err := s.contestantRepository.JoinContest(*d)
	if err != nil {
		return fmt.Errorf("failed to join contest: %w", err)
	}
	return nil
}
