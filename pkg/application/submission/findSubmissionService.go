package submission

import (
	"errors"
	"fmt"
	"sort"

	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type FindSubmissionService struct {
	submissionRepository repository.SubmissionRepository
}

func NewFindSubmissionService(submissionRepository repository.SubmissionRepository) *FindSubmissionService {
	return &FindSubmissionService{submissionRepository: submissionRepository}
}

func (s FindSubmissionService) FindByID(id id.SnowFlakeID) (*Data, error) {
	su, err := s.submissionRepository.FindSubmissionByID(id)
	if err != nil {
		return nil, err
	}
	return DomainToData(*su), nil
}

func (s FindSubmissionService) FindTask() (*Data, error) {
	res, err := s.submissionRepository.FindSubmissionByStatus("WE")
	if err != nil {
		return nil, fmt.Errorf("failed to find task: %w", err)
	}

	if len(res) == 0 {
		return nil, errors.New("failed to find task: not found")
	}
	// 早いものからソートする
	sort.Slice(res, func(i, j int) bool {
		return res[i].GetSubmittedAt().Before(res[j].GetSubmittedAt())
	})

	task := res[0]
	res[0].SetResult("WJ")
	if _, err := s.submissionRepository.UpdateSubmissionResult(res[0]); err != nil {
		return nil, fmt.Errorf("failed to update submission status: %w", err)
	}

	return DomainToData(task), nil
}
