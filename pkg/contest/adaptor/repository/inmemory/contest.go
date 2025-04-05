package inmemory

import (
	"errors"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ContestRepository struct {
	data []model.Contest
}

func (c *ContestRepository) CreateContest(d model.Contest) error {
	c.data = append(c.data, d)
	return nil
}

func (c *ContestRepository) FindAllContests() ([]model.Contest, error) {
	return c.data, nil
}

func (c *ContestRepository) FindContestByID(id id.SnowFlakeID) (*model.Contest, error) {
	for _, v := range c.data {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func (c *ContestRepository) FindContestByTitle(title string) (*model.Contest, error) {
	for _, v := range c.data {
		if v.GetTitle() == title {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func NewContestRepository(d []model.Contest) *ContestRepository {
	return &ContestRepository{data: d}
}
