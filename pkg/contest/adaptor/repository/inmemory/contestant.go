package inmemory

import (
	"errors"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ContestantRepository struct {
	data []model.Contestant
}

func NewContestantRepository(d []model.Contestant) *ContestantRepository {
	return &ContestantRepository{data: d}
}

func (c *ContestantRepository) JoinContest(d model.Contestant) error {
	c.data = append(c.data, d)
	return nil
}

func (c *ContestantRepository) FindContestantByID(id id.SnowFlakeID) (*model.Contestant, error) {
	for _, v := range c.data {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func (c *ContestantRepository) FindContestantByUserID(id id.SnowFlakeID) ([]model.Contestant, error) {
	res := make([]model.Contestant, 0)
	for _, v := range c.data {
		if v.GetUserID() == id {
			res = append(res, v)
		}
	}
	return res, nil
}

func (c *ContestantRepository) FindContestantByContestID(id id.SnowFlakeID) ([]model.Contestant, error) {
	res := make([]model.Contestant, 0)
	for _, v := range c.data {
		if v.GetContestID() == id {
			res = append(res, v)
		}
	}
	return res, nil
}
