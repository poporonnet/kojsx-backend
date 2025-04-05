package repository

import (
	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ContestRepository interface {
	// CreateContest コンテストを作成
	CreateContest(d domain.Contest) error
	// FindAllContests コンテストをすべて取得
	FindAllContests() ([]domain.Contest, error)
	// FindContestByID コンテストを1つ取得(IDで検索
	FindContestByID(id id.SnowFlakeID) (*domain.Contest, error)
	// FindContestByTitle コンテストを1つ取得(タイトルで検索
	FindContestByTitle(title string) (*domain.Contest, error)
}
