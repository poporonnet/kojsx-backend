package repository

import (
	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type SubmissionRepository interface {
	// CreateSubmission 提出を作成
	CreateSubmission(submission domain.Submission) error
	// FindSubmissionByID 提出を1つ取得
	FindSubmissionByID(id id.SnowFlakeID) (*domain.Submission, error)
	// FindSubmissionByStatus 提出をステータスで検索
	FindSubmissionByStatus(status string) ([]domain.Submission, error)
	// FindSubmissionByProblemID 問題IDで提出を検索
	FindSubmissionByProblemID(id id.SnowFlakeID) ([]domain.Submission, error)
	// UpdateSubmissionResult 提出を更新
	UpdateSubmissionResult(submission domain.Submission) (*domain.Submission, error)
}
