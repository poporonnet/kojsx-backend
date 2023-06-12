package inmemory

import (
	"errors"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type SubmissionRepository struct {
	submission []domain.Submission
}

func (s *SubmissionRepository) UpdateSubmissionResult(submission domain.Submission) (*domain.Submission, error) {
	for i, v := range s.submission {
		if v.GetID() == submission.GetID() {
			s.submission[i] = submission
			return &submission, nil
		}
	}

	return nil, errors.New("not found")
}

func (s *SubmissionRepository) CreateSubmission(submission domain.Submission) error {
	s.submission = append(s.submission, submission)
	return nil
}

func NewSubmissionRepository(s []domain.Submission) *SubmissionRepository {
	return &SubmissionRepository{submission: s}
}

func (s *SubmissionRepository) FindSubmissionByID(id id.SnowFlakeID) (*domain.Submission, error) {
	for _, v := range s.submission {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("no such submission")
}

func (s *SubmissionRepository) FindSubmissionByStatus(status string) ([]domain.Submission, error) {
	res := make([]domain.Submission, 0)
	for _, v := range s.submission {
		if v.GetResult() == status {
			res = append(res, v)
		}
	}
	return res, nil
}
