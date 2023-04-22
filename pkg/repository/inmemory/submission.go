package inmemory

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type SubmissionRepository struct {
	submission []domain.Submission
	result     []domain.SubmissionResult
}

func NewSubmissionRepository(s []domain.Submission, r []domain.SubmissionResult) *SubmissionRepository {
	return &SubmissionRepository{submission: s, result: r}
}

func (s SubmissionRepository) FindSubmissionByID(id id.SnowFlakeID) *domain.Submission {
	for _, v := range s.submission {
		if v.GetID() == id {
			return &v
		}
	}
	return nil
}
