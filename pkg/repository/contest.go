package repository

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ContestRepository interface {
	// CreateContest コンテストを作成
	CreateContest(d domain.Contest) error
	// FindAllContests コンテストをすべて取得
	FindAllContests() []domain.Contest
	// FindContestByID コンテストを1つ取得(IDで検索
	FindContestByID(id id.SnowFlakeID) *domain.Contest
	// FindContestByTitle コンテストを1つ取得(タイトルで検索
	FindContestByTitle(title string) *domain.Contest
}
