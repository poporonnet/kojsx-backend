package contest

import (
	"errors"
	"fmt"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/domain/service"
	"github.com/poporonnet/kojsx-backend/pkg/repository"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type CreateContestService struct {
	contestRepository  repository.ContestRepository
	contestService     *service.ContestService
	joinContestService JoinContestService
}

func NewCreateContestService(
	contestRepository repository.ContestRepository,
	contestantRepository repository.ContestantRepository,
	contestantService service.ContestantService,
) *CreateContestService {
	return &CreateContestService{
		contestRepository:  contestRepository,
		contestService:     service.NewContestService(contestRepository),
		joinContestService: *NewJoinContestService(contestantRepository, contestantService),
	}
}

func (s *CreateContestService) Handle(args CreateContestArgs) (*Data, error) {
	gen := id.NewSnowFlakeIDGenerator()
	i := gen.NewID(time.Now())
	c := domain.NewContest(i)

	// コンテストの作成は管理者のみできる
	if !args.User.IsAdmin() {
		return nil, errors.New("can't create contest: forbidden")
	}

	if err := c.SetTitle(args.Title); err != nil {
		return nil, fmt.Errorf("failed to set title: %w", err)
	}
	if err := c.SetDescription(args.Description); err != nil {
		return nil, fmt.Errorf("failed to set description: %w", err)
	}
	if err := c.SetStartAt(args.StartAt); err != nil {
		return nil, fmt.Errorf("failed to set startAt: %w", err)
	}
	if err := c.SetEndAt(args.EndAt); err != nil {
		return nil, fmt.Errorf("failed to set endAt: %w", err)
	}

	if s.contestService.IsExists(*c) {
		return nil, errors.New("AlreadyExists")
	}

	if err := s.contestRepository.CreateContest(*c); err != nil {
		return nil, fmt.Errorf("failed to create contest: %w", err)
	}
	time.Sleep(1 * time.Millisecond)
	// コンテストの作成者はコンテストの管理者になる
	err := s.joinContestService.Join(i, args.User, domain.ContestAdmin)
	if err != nil {
		return nil, err
	}
	r := DomainToData(*c)
	return &r, nil
}

type CreateContestArgs struct {
	// Title コンテストのタイトル
	Title string
	// Description コンテストの説明
	Description string
	// StartAt 開始時刻
	StartAt time.Time
	// EndAt 終了時刻
	EndAt time.Time
	// User 操作を行うユーザー
	User domain.User
}
