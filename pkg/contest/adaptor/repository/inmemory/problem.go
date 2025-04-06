package inmemory

import (
	"errors"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ProblemRepository struct {
	data []model.Problem
}

func NewProblemRepository(data []model.Problem) *ProblemRepository {
	return &ProblemRepository{data: data}
}

func (p *ProblemRepository) FindProblemByContestID(id id.SnowFlakeID) ([]model.Problem, error) {
	res := make([]model.Problem, 0)
	for _, v := range p.data {
		if v.GetContestID() == id {
			res = append(res, v)
		}
	}
	return res, nil
}

func (p *ProblemRepository) CreateProblem(in model.Problem) error {
	p.data = append(p.data, in)
	return nil
}

func (p *ProblemRepository) FindProblemByID(id id.SnowFlakeID) (*model.Problem, error) {
	for _, v := range p.data {
		if v.GetProblemID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func (p *ProblemRepository) FindProblemByTitle(name string) (*model.Problem, error) {
	for _, v := range p.data {
		if v.GetTitle() == name {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func (p *ProblemRepository) FindCaseSetByID(id id.SnowFlakeID) (*model.Caseset, error) {
	for _, k := range p.data {
		for _, v := range k.GetCaseSets() {
			if v.GetID() == id {
				return &v, nil
			}
		}
	}
	return nil, errors.New("not found")
}

func (p *ProblemRepository) FindCaseByID(id id.SnowFlakeID) (*model.Case, error) {
	for _, k := range p.data {
		for _, j := range k.GetCaseSets() {
			for _, v := range j.GetCases() {
				if v.GetID() == id {
					return &v, nil
				}
			}
		}
	}
	return nil, errors.New("not found")
}
