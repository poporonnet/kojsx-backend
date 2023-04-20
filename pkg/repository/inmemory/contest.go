package inmemory

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ContestRepository struct {
	data []domain.Contest
}

func (c ContestRepository) FindAllContests() []domain.Contest {
	return c.data
}

func (c ContestRepository) FindContestByID(id id.SnowFlakeID) *domain.Contest {
	for _, v := range c.data {
		if v.GetID() == id {
			return &v
		}
	}
	return nil
}

func (c ContestRepository) FindContestByTitle(title string) *domain.Contest {
	for _, v := range c.data {
		if v.GetTitle() == title {
			return &v
		}
	}
	return nil
}

func NewContestRepository(d []domain.Contest) *ContestRepository {
	return &ContestRepository{data: d}
}
