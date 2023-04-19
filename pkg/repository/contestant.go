package repository

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ContestantRepository interface {
	// FindContestantByID 参加者を1つ取得(参加者IDで検索
	FindContestantByID(id id.SnowFlakeID) *domain.Contestant
	// FindContestantByUserID 参加者を取得(ユーザーIDで検索
	FindContestantByUserID(id id.SnowFlakeID) []domain.Contestant
	// FindContestantByContestID 参加者を取得(コンテストIDで検索
	FindContestantByContestID(id id.SnowFlakeID) []domain.Contestant
}
