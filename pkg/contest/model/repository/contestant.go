package repository

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ContestantRepository interface {
	// JoinContest コンテストに参加
	JoinContest(d model.Contestant) error
	// FindContestantByID 参加者を1つ取得(参加者IDで検索
	FindContestantByID(id id.SnowFlakeID) (*model.Contestant, error)
	// FindContestantByUserID 参加者を取得(ユーザーIDで検索
	FindContestantByUserID(id id.SnowFlakeID) ([]model.Contestant, error)
	// FindContestantByContestID 参加者を取得(コンテストIDで検索
	FindContestantByContestID(id id.SnowFlakeID) ([]model.Contestant, error)
}
