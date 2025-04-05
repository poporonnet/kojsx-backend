package mongodb

import (
	"context"
	"fmt"

	"github.com/poporonnet/kojsx-backend/pkg/utils"

	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/repository/mongodb/entity"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
	"go.mongodb.org/mongo-driver/bson"
)

type ContestRepository struct {
	client Client
}

func (c ContestRepository) CreateContest(d domain.Contest) error {
	e := entity.Contest{
		ID:          d.GetID(),
		Title:       d.GetTitle(),
		Description: d.GetDescription(),
		StartAt:     d.GetStartAt(),
		EndAt:       d.GetEndAt(),
	}

	_, err := c.client.Cli.Database("kojs").Collection("contest").InsertOne(context.Background(), e)
	if err != nil {
		utils.SugarLogger.Errorf("failed to create contest: %v", err)
		return fmt.Errorf("failed to create contest: %w", err)
	}

	return nil
}

func (c ContestRepository) FindAllContests() ([]domain.Contest, error) {
	filter := &bson.D{}
	cursor, err := c.client.Cli.Database("kojs").Collection("contest").Find(context.Background(), filter)
	if err != nil {
		utils.SugarLogger.Errorf("failed to find contests: %v", err)
		return []domain.Contest{}, fmt.Errorf("failed to find contests: %w", err)
	}

	var contest []entity.Contest
	if err := cursor.All(context.Background(), &contest); err != nil {
		utils.SugarLogger.Errorf("failed to get value from cursor: %v", err)
		return []domain.Contest{}, fmt.Errorf("failed get value from cursor: %w", err)
	}

	res := make([]domain.Contest, len(contest))
	for i, v := range contest {
		res[i] = v.ToDomain()
	}

	return res, nil
}

func (c ContestRepository) FindContestByID(id id.SnowFlakeID) (*domain.Contest, error) {
	filter := &bson.M{"_id": id}
	result := c.client.Cli.Database("kojs").Collection("contest").FindOne(context.Background(), filter)

	var contest entity.Contest
	if err := result.Decode(&contest); err != nil {
		utils.SugarLogger.Errorf("failed to decode contest data: %v", err)
		return nil, fmt.Errorf("failed to decode contest data: %w", err)
	}
	res := contest.ToDomain()
	return &res, nil
}

func (c ContestRepository) FindContestByTitle(title string) (*domain.Contest, error) {
	filter := &bson.M{"title": title}
	result := c.client.Cli.Database("kojs").Collection("contest").FindOne(context.Background(), filter)

	var contest entity.Contest
	if err := result.Decode(&contest); err != nil {
		utils.SugarLogger.Errorf("failed to decode contest data: %v", err)
		return nil, fmt.Errorf("failed to decode contest data: %w", err)
	}
	res := contest.ToDomain()
	return &res, nil
}

func NewContestRepository(cli Client) *ContestRepository {
	return &ContestRepository{client: cli}
}
