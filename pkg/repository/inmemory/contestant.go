package inmemory

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ContestantRepository struct {
	data []domain.Contestant
}

func NewContestantRepository(d []domain.Contestant) *ContestantRepository {
	return &ContestantRepository{data: d}
}

func (c *ContestantRepository) JoinContest(d domain.Contestant) error {
	c.data = append(c.data, d)
	return nil
}

func (c *ContestantRepository) FindContestantByID(id id.SnowFlakeID) *domain.Contestant {
	for _, v := range c.data {
		if v.GetID() == id {
			return &v
		}
	}
	return nil
}

func (c *ContestantRepository) FindContestantByUserID(id id.SnowFlakeID) []domain.Contestant {
	res := make([]domain.Contestant, 0)
	for _, v := range c.data {
		if v.GetUserID() == id {
			res = append(res, v)
		}
	}
	return res
}

func (c *ContestantRepository) FindContestantByContestID(id id.SnowFlakeID) []domain.Contestant {
	res := make([]domain.Contestant, 0)
	for _, v := range c.data {
		if v.GetContestID() == id {
			res = append(res, v)
		}
	}
	return res
}
