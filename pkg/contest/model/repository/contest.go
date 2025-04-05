package repository

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ContestRepository interface {
	// CreateContest コンテストを作成
	CreateContest(d model.Contest) error
	// FindAllContests コンテストをすべて取得
	FindAllContests() ([]model.Contest, error)
	// FindContestByID コンテストを1つ取得(IDで検索
	FindContestByID(id id.SnowFlakeID) (*model.Contest, error)
	// FindContestByTitle コンテストを1つ取得(タイトルで検索
	FindContestByTitle(title string) (*model.Contest, error)
}
