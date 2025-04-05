package contest_test

import (
	"testing"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/domain/service"

	"github.com/poporonnet/kojsx-backend/pkg/application/contest"
	"github.com/poporonnet/kojsx-backend/pkg/utils/seed"

	"github.com/poporonnet/kojsx-backend/pkg/repository/inmemory"
	"github.com/stretchr/testify/assert"
)

func TestCreateContestService_Handle(t *testing.T) {
	r := inmemory.NewContestRepository(seed.NewSeeds().Contests)
	participant := inmemory.NewContestantRepository([]domain.Contestant{})
	cContestService := contest.NewCreateContestService(r, participant, *service.NewContestantService(participant))

	t.Run("重複がなければ成功する", func(t *testing.T) {
		_, err := cContestService.Handle(
			contest.CreateContestArgs{
				Title:       "Test Contest",
				Description: "hello world",
				StartAt:     time.Now(),
				EndAt:       time.Now().Add(60 * time.Minute),
			},
		)
		assert.Equal(t, nil, err)
	})

	t.Run("重複があれば失敗する", func(t *testing.T) {
		_, err2 := cContestService.Handle(
			contest.CreateContestArgs{
				Title:       "Test Contest",
				Description: "hello world",
				StartAt:     time.Now(),
				EndAt:       time.Now().Add(60 * time.Minute),
			},
		)
		assert.NotEqual(t, nil, err2)
	})

	t.Run("通常ユーザーはコンテストの作成ができない", func(t *testing.T) {
		_, err := cContestService.Handle(
			contest.CreateContestArgs{
				Title:       "Test Contest 2",
				Description: "hello world",
				StartAt:     time.Now(),
				EndAt:       time.Now().Add(60 * time.Minute),
				User:        seed.NewSeeds().Users[1],
			},
		)
		assert.NotEqual(t, nil, err)
	})

	t.Run("システムの管理者のみコンテストの作成ができる", func(t *testing.T) {
		time.Sleep(1 * time.Millisecond)
		_, err := cContestService.Handle(
			contest.CreateContestArgs{
				Title:       "Test Contest 3",
				Description: "hello world",
				StartAt:     time.Now(),
				EndAt:       time.Now().Add(60 * time.Minute),
				User:        seed.NewSeeds().Users[0],
			},
		)
		assert.Equal(t, nil, err)
	})

	t.Run("コンテストの作成者はコンテストの管理者になる", func(t *testing.T) {
		time.Sleep(1 * time.Millisecond)
		cont, _ := cContestService.Handle(
			contest.CreateContestArgs{
				Title:       "Test Contest 4",
				Description: "hello world",
				StartAt:     time.Now(),
				EndAt:       time.Now().Add(60 * time.Minute),
				User:        seed.NewSeeds().Users[0],
			},
		)
		res, _ := participant.FindContestantByUserID(seed.NewSeeds().Users[0].GetID())
		co := &domain.Contestant{}
		for _, v := range res {
			if v.GetContestID() == cont.GetID() {
				co = &v
			}
		}

		assert.Equal(t, true, co.IsAdmin())
	})

}
