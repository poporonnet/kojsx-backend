package submission

import (
	"errors"
	"fmt"
	"sort"
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/utils"

	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/domain/service"
	"github.com/poporonnet/kojsx-backend/pkg/repository"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type CreateSubmissionService struct {
	repository        repository.SubmissionRepository
	problemRepository repository.ProblemRepository
	service           service.SubmissionService
	idGenerator       id.Generator
}

func NewCreateSubmissionService(
	repo repository.SubmissionRepository,
	service service.SubmissionService,
	problemRepository repository.ProblemRepository,
) *CreateSubmissionService {
	return &CreateSubmissionService{
		repo,
		problemRepository,
		service,
		id.NewSnowFlakeIDGenerator(),
	}
}

func (s CreateSubmissionService) Handle(pID, cID id.SnowFlakeID, lang, code string) (*domain.Submission, error) {
	newID := s.idGenerator.NewID(time.Now())

	d, err := domain.NewSubmission(newID, pID, cID, lang, code, time.Now())
	if err != nil {
		return nil, fmt.Errorf("failed to create submission: %w", err)
	}

	if s.service.IsExists(*d) {
		return nil, errors.New("submission exists")
	}

	if err := s.repository.CreateSubmission(*d); err != nil {
		return nil, fmt.Errorf("failed to create submission: %w", err)
	}

	return d, nil
}

type CreateResultArgs struct {
	Result     string
	Output     string
	CaseName   string
	ExitStatus int
	ExecTime   int
	ExecMemory int
}

func (s CreateSubmissionService) CreateResult(submissionID id.SnowFlakeID, args []CreateResultArgs) (Data, error) {
	submission, err := s.repository.FindSubmissionByID(submissionID)
	if err != nil {
		return Data{}, err
	}
	problem, err := s.problemRepository.FindProblemByID(submission.GetProblemID())
	if err != nil {
		return Data{}, err
	}
	results := make([]domain.SubmissionResult, len(args))
	for i, v := range args {
		time.Sleep(2 * time.Millisecond)
		d := domain.NewSubmissionResult(
			s.idGenerator.NewID(time.Now()),
			v.Result,
			v.Output,
			v.CaseName,
			v.ExitStatus,
			v.ExecTime,
			v.ExecMemory,
		)
		results[i] = *d
	}

	scoreResult, err := scoring(*problem, results)
	if err != nil {
		return Data{}, err
	}
	_ = submission.SetPoint(scoreResult.Point)
	submission.SetResult(scoreResult.Status)
	for _, v := range scoreResult.SubmissionResult {
		err = submission.AddResult(v)
		if err != nil {
			utils.Logger.Sugar().Warnf("%v", err)
		}
	}

	sort.Slice(scoreResult.SubmissionResult, func(i, j int) bool {
		return scoreResult.SubmissionResult[i].GetExecTime() > scoreResult.SubmissionResult[j].GetExecTime()
	})
	sort.Slice(scoreResult.SubmissionResult, func(i, j int) bool {
		return scoreResult.SubmissionResult[i].GetExecMemory() > scoreResult.SubmissionResult[j].GetExecMemory()
	})
	submission.SetExecTime(scoreResult.SubmissionResult[0].GetExecTime())
	submission.SetExecMemory(scoreResult.SubmissionResult[0].GetExecMemory())
	_, err = s.repository.UpdateSubmissionResult(*submission)
	if err != nil {
		return Data{}, err
	}

	ret := *DomainToData(*submission)
	return ret, nil
}
