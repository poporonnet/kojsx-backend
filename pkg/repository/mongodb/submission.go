package mongodb

import (
	"context"
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/utils"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/mongodb/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"go.mongodb.org/mongo-driver/bson"
)

type SubmissionRepository struct {
	client Client
}

func (s SubmissionRepository) CreateSubmission(submission domain.Submission) error {
	e := entity.Submission{
		ID:           submission.GetID(),
		ProblemID:    submission.GetProblemID(),
		ContestantID: submission.GetContestantID(),
		Point:        submission.GetPoint(),
		Lang:         submission.GetLang(),
		CodeLength:   submission.GetCodeLength(),
		Result:       submission.GetResult(),
		ExecTime:     submission.GetExecTime(),
		ExecMemory:   submission.GetExecMemory(),
		Code:         submission.GetCode(),
		SubmittedAt:  submission.GetSubmittedAt(),
		Results:      nil,
	}

	_, err := s.client.Cli.Database("kojs").Collection("submission").InsertOne(context.Background(), e)
	if err != nil {
		utils.SugarLogger.Errorf("failed to create submision: %v", err)
		return fmt.Errorf("failed to create submission: %w", err)
	}

	return nil
}

func (s SubmissionRepository) FindSubmissionByID(id id.SnowFlakeID) (*domain.Submission, error) {
	filter := &bson.M{"_id": id}

	result := s.client.Cli.Database("kojs").Collection("submission").FindOne(context.Background(), filter)

	var submission entity.Submission
	if err := result.Decode(&submission); err != nil {
		utils.SugarLogger.Errorf("failed to decode submission data: %v", err)
		return nil, fmt.Errorf("failed to decode submission data: %w", err)
	}

	res := submission.ToDomain()
	return &res, nil
}

func (s SubmissionRepository) FindSubmissionByStatus(status string) ([]domain.Submission, error) {
	filter := &bson.M{"result": status}

	cursor, err := s.client.Cli.Database("kojs").Collection("submission").Find(context.Background(), filter)
	if err != nil {
		utils.SugarLogger.Errorf("failed to find submission by status: %v", err)
		return nil, fmt.Errorf("failed to find submission by status: %w", err)
	}

	var submission []entity.Submission
	if err := cursor.All(context.Background(), &submission); err != nil {
		utils.SugarLogger.Errorf("failed to decode submission data: %v", err)
		return nil, fmt.Errorf("failed to decode submission data: %w", err)
	}

	res := make([]domain.Submission, len(submission))
	for i, v := range submission {
		res[i] = v.ToDomain()
	}
	return res, nil
}

func (s SubmissionRepository) UpdateSubmissionResult(submission domain.Submission) (*domain.Submission, error) {
	r := make([]entity.SubmissionResult, len(submission.GetResults()))
	for i, v := range submission.GetResults() {
		r[i] = entity.SubmissionResult{
			ID:         v.GetID(),
			Result:     v.GetResult(),
			Output:     v.GetOutput(),
			CaseName:   v.GetCaseName(),
			ExitStatus: v.GetExitStatus(),
			ExecTime:   v.GetExecTime(),
			ExecMemory: v.GetExecMemory(),
		}
	}
	e := entity.Submission{
		ID:           submission.GetID(),
		ProblemID:    submission.GetProblemID(),
		ContestantID: submission.GetContestantID(),
		Point:        submission.GetPoint(),
		Lang:         submission.GetLang(),
		CodeLength:   submission.GetCodeLength(),
		Result:       submission.GetResult(),
		ExecTime:     submission.GetExecTime(),
		ExecMemory:   submission.GetExecMemory(),
		Code:         submission.GetCode(),
		SubmittedAt:  submission.GetSubmittedAt(),
		Results:      r,
	}
	_, err := s.client.Cli.Database("kojs").Collection("submission").ReplaceOne(context.Background(), bson.M{"_id": submission.GetID()}, e)
	if err != nil {
		utils.SugarLogger.Errorf("failed to update submission: %v", err)
		return nil, fmt.Errorf("failed to update submission: %w", err)
	}

	return &submission, nil
}

func NewSubmissionRepository(client Client) *SubmissionRepository {
	return &SubmissionRepository{client: client}
}
