package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/mct-joken/kojs5-backend/pkg/utils"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/mongodb/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"go.mongodb.org/mongo-driver/bson"
)

type ProblemRepository struct {
	cli Client
}

func (p ProblemRepository) CreateProblem(in domain.Problem) error {
	sets := in.GetCaseSets()
	setsEntity := make([]entity.CaseSet, len(sets))
	for i, v := range sets {
		caseEntity := make([]entity.Case, len(v.GetCases()))
		for j, k := range v.GetCases() {
			caseEntity[j] = entity.Case{
				ID:        k.GetID(),
				CaseSetID: k.GetCasesetID(),
				In:        k.GetIn(),
				Out:       k.GetOut(),
			}
		}

		setsEntity[i] = entity.CaseSet{
			ID:    v.GetID(),
			Name:  v.GetName(),
			Point: v.GetPoint(),
			Cases: caseEntity,
		}
	}
	e := entity.Problem{
		ID:          in.GetProblemID(),
		ContestID:   in.GetContestID(),
		Index:       in.GetIndex(),
		Title:       in.GetTitle(),
		Text:        in.GetText(),
		Point:       in.GetPoint(),
		MemoryLimit: in.GetMemoryLimit(),
		TimeLimit:   in.GetTimeLimit(),
		CaseSets:    setsEntity,
	}

	_, err := p.cli.Cli.Database("kojs").Collection("problem").InsertOne(context.Background(), e)
	if err != nil {
		utils.SugarLogger.Errorf("failed to create problem: %v", err)
		return fmt.Errorf("failed to create problem: %w", err)
	}

	return nil
}

func (p ProblemRepository) FindProblemByID(id id.SnowFlakeID) (*domain.Problem, error) {
	result := p.cli.Cli.Database("kojs").Collection("problem").FindOne(context.Background(), &bson.M{"_id": id})

	var problem entity.Problem
	if err := result.Decode(&problem); err != nil {
		utils.SugarLogger.Errorf("failed to decode problem data: %v", err)
		return nil, fmt.Errorf("failed to decode problem data: %w", err)
	}
	res := problem.ToDomain()
	return &res, nil
}

func (p ProblemRepository) FindProblemByTitle(name string) (*domain.Problem, error) {
	result := p.cli.Cli.Database("kojs").Collection("problem").FindOne(context.Background(), &bson.M{"name": name})

	var problem entity.Problem
	if err := result.Decode(&problem); err != nil {
		utils.SugarLogger.Errorf("failed to decode problem data: %v", err)
		return nil, fmt.Errorf("failed to decode problem data: %w", err)
	}
	res := problem.ToDomain()
	return &res, nil
}

func (p ProblemRepository) FindCaseSetByID(id id.SnowFlakeID) (*domain.Caseset, error) {
	filter := &bson.M{"casesets.id": id}
	cursor := p.cli.Cli.Database("kojs").Collection("problem").FindOne(context.Background(), filter)

	var problem entity.Problem
	if err := cursor.Decode(&problem); err != nil {
		utils.SugarLogger.Errorf("failed to decode problem data: %v", err)
		return nil, fmt.Errorf("failed to decode problem data: %w", err)
	}
	res := problem.ToDomain()
	for _, v := range res.GetCaseSets() {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("no such case set")
}

func (p ProblemRepository) FindCaseByID(id id.SnowFlakeID) (*domain.Case, error) {
	cursor := p.cli.Cli.Database("kojs").Collection("problem").FindOne(context.Background(), &bson.M{"casesets.cases.id": id})

	var problem entity.Problem
	if err := cursor.Decode(&problem); err != nil {
		utils.SugarLogger.Errorf("failed to decode problem data: %v", err)
		return nil, fmt.Errorf("failed to decode problem data: %w", err)
	}
	res := problem.ToDomain()
	for _, v := range res.GetCaseSets() {
		for _, k := range v.GetCases() {
			if k.GetID() == id {
				return &k, nil
			}
		}
	}

	return nil, errors.New("no such case")
}

func (p ProblemRepository) FindProblemByContestID(id id.SnowFlakeID) ([]domain.Problem, error) {
	cursor, err := p.cli.Cli.Database("kojs").Collection("problem").Find(context.Background(), &bson.M{"contestID": id})
	if err != nil {
		utils.SugarLogger.Errorf("failed to find problems: %v", err)
		return []domain.Problem{}, fmt.Errorf("failed to find problems: %w", err)
	}

	var problem []entity.Problem
	if err := cursor.All(context.Background(), &problem); err != nil {
		utils.SugarLogger.Errorf("failed to decode problems: %v", err)
		return []domain.Problem{}, fmt.Errorf("failed to decode problems: %w", err)
	}
	res := make([]domain.Problem, len(problem))
	for i, v := range problem {
		res[i] = v.ToDomain()
	}

	return res, nil
}

func NewProblemRepository(cli Client) *ProblemRepository {
	return &ProblemRepository{cli: cli}
}
