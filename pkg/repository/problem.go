package repository

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ProblemRepository interface {
	CreateProblem(in domain.Problem) error
	// FindProblemByContestID コンテストの問題リストを取得
	FindProblemByContestID(id id.SnowFlakeID) []domain.Problem
	// FindProblemByID 問題を1つ取得
	FindProblemByID(id id.SnowFlakeID) *domain.Problem
	// FindProblemByTitle 問題をタイトルで取得
	FindProblemByTitle(name string) *domain.Problem

	// FindCaseSetByID ケースセットを1つ取得
	FindCaseSetByID(id id.SnowFlakeID) *domain.Caseset

	// FindCaseByID ケースを1つ取得
	FindCaseByID(id id.SnowFlakeID) *domain.Case
}
