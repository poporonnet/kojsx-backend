package submission

import (
	"errors"
	"fmt"
	"sort"

	"github.com/mct-joken/kojs5-backend/pkg/utils"

	"github.com/mct-joken/kojs5-backend/pkg/application/problem"
	"github.com/mct-joken/kojs5-backend/pkg/domain"

	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type FindSubmissionService struct {
	submissionRepository repository.SubmissionRepository
	problemRepository    repository.ProblemRepository
}

func NewFindSubmissionService(submissionRepository repository.SubmissionRepository, problemRepository repository.ProblemRepository) *FindSubmissionService {
	return &FindSubmissionService{submissionRepository: submissionRepository, problemRepository: problemRepository}
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
		utils.SugarLogger.Errorf("failed to update submission status: %v", err)
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
		utils.SugarLogger.Errorf("failed to update submission status: %v", err)
		return nil, fmt.Errorf("failed to update submission status: %w", err)
	}

	return DomainToData(task), nil
}

func (s FindSubmissionService) FindByContestID(i id.SnowFlakeID) (FindByContestIDResult, error) {
	// コンテストの問題リストを取得
	problemList, err := s.problemRepository.FindProblemByContestID(i)
	if err != nil {
		return FindByContestIDResult{}, err
	}

	// subs 提出
	subs := make([]domain.Submission, 0)
	for _, v := range problemList {
		// 問題リストで提出を検索
		su, err := s.submissionRepository.FindSubmissionByProblemID(v.GetProblemID())
		if err != nil {
			return FindByContestIDResult{}, err
		}
		subs = append(subs, su...)
	}

	sd := make([]Data, len(subs))
	for j, k := range subs {
		sd[j] = *DomainToData(k)
	}

	pd := make([]problem.Data, len(problemList))
	for j, k := range problemList {
		pd[j] = problem.DomainToData(k)
	}

	return FindByContestIDResult{
		S: sd,
		P: pd,
	}, nil
}

type FindByContestIDResult struct {
	S []Data
	P []problem.Data
}
