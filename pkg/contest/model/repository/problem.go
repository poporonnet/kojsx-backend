package repository

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ProblemRepository interface {
	CreateProblem(in model.Problem) error
	// FindProblemByContestID コンテストの問題リストを取得
	FindProblemByContestID(id id.SnowFlakeID) ([]model.Problem, error)
	// FindProblemByID 問題を1つ取得
	FindProblemByID(id id.SnowFlakeID) (*model.Problem, error)
	// FindProblemByTitle 問題をタイトルで取得
	FindProblemByTitle(name string) (*model.Problem, error)

	// FindCaseSetByID ケースセットを1つ取得
	FindCaseSetByID(id id.SnowFlakeID) (*model.Caseset, error)

	// FindCaseByID ケースを1つ取得
	FindCaseByID(id id.SnowFlakeID) (*model.Case, error)
}
