package inmemory

import (
	"errors"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ContestRepository struct {
	data []domain.Contest
}

func (c *ContestRepository) CreateContest(d domain.Contest) error {
	c.data = append(c.data, d)
	return nil
}

func (c *ContestRepository) FindAllContests() ([]domain.Contest, error) {
	return c.data, nil
}

func (c *ContestRepository) FindContestByID(id id.SnowFlakeID) (*domain.Contest, error) {
	for _, v := range c.data {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func (c *ContestRepository) FindContestByTitle(title string) (*domain.Contest, error) {
	for _, v := range c.data {
		if v.GetTitle() == title {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func NewContestRepository(d []domain.Contest) *ContestRepository {
	return &ContestRepository{data: d}
}
