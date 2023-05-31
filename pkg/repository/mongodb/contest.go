package mongodb

import (
	"context"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/repository/mongodb/entity"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
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

	_, err := c.client.Cli.Database("kojs").Collection("user").InsertOne(context.Background(), e)
	if err != nil {
		return err
	}

	return nil
}

func (c ContestRepository) FindAllContests() []domain.Contest {
	filter := &bson.D{}
	cursor, err := c.client.Cli.Database("kojs").Collection("user").Find(context.Background(), filter)
	if err != nil {
		return []domain.Contest{}
	}

	var contest []entity.Contest
	if err := cursor.All(context.Background(), &contest); err != nil {
		return []domain.Contest{}
	}

	res := make([]domain.Contest, len(contest))
	for i, v := range contest {
		res[i] = v.ToDomain()
	}

	return res
}

func (c ContestRepository) FindContestByID(id id.SnowFlakeID) *domain.Contest {
	filter := &bson.M{"_id": id}
	result := c.client.Cli.Database("kojs").Collection("user").FindOne(context.Background(), filter)

	var contest entity.Contest
	if err := result.Decode(&contest); err != nil {
		return nil
	}
	res := contest.ToDomain()
	return &res
}

func (c ContestRepository) FindContestByTitle(title string) *domain.Contest {
	filter := &bson.M{"title": title}
	result := c.client.Cli.Database("kojs").Collection("user").FindOne(context.Background(), filter)

	var contest entity.Contest
	if err := result.Decode(&contest); err != nil {
		return nil
	}
	res := contest.ToDomain()
	return &res
}

func NewContestRepository(cli Client) *ContestRepository {
	return &ContestRepository{client: cli}
}
