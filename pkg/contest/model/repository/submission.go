package repository

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type SubmissionRepository interface {
	// CreateSubmission 提出を作成
	CreateSubmission(submission model.Submission) error
	// FindSubmissionByID 提出を1つ取得
	FindSubmissionByID(id id.SnowFlakeID) (*model.Submission, error)
	// FindSubmissionByStatus 提出をステータスで検索
	FindSubmissionByStatus(status string) ([]model.Submission, error)
	// FindSubmissionByProblemID 問題IDで提出を検索
	FindSubmissionByProblemID(id id.SnowFlakeID) ([]model.Submission, error)
	// UpdateSubmissionResult 提出を更新
	UpdateSubmissionResult(submission model.Submission) (*model.Submission, error)
}
