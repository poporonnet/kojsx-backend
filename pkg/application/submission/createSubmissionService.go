package submission

import (
	"errors"
	"fmt"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/domain/service"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
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

func (s CreateSubmissionService) CreateResult(submissionID id.SnowFlakeID, args []CreateResultArgs) error {
	newID := s.idGenerator.NewID(time.Now())
	submission, err := s.repository.FindSubmissionByID(submissionID)
	if err != nil {
		return err
	}
	problem, err := s.problemRepository.FindProblemByID(submission.GetProblemID())
	if err != nil {
		return err
	}

	results := make([]domain.SubmissionResult, len(args))
	for i, v := range args {
		d := domain.NewSubmissionResult(newID, v.Result, v.Output, v.CaseName, v.ExitStatus, v.ExecTime, v.ExecMemory)
		results[i] = *d
	}
	scoreResult, _ := scoring(*problem, results)
	results = scoreResult.SubmissionResult

	_ = submission.SetPoint(scoreResult.Point)
	submission.SetResult(scoreResult.Status)
	_, err = s.repository.UpdateSubmissionResult(*submission)
	if err != nil {
		return err
	}
	return nil
}
