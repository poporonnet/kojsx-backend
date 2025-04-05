package inmemory

import (
	"errors"

	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type ProblemRepository struct {
	data []domain.Problem
}

func NewProblemRepository(data []domain.Problem) *ProblemRepository {
	return &ProblemRepository{data: data}
}

func (p *ProblemRepository) FindProblemByContestID(id id.SnowFlakeID) ([]domain.Problem, error) {
	res := make([]domain.Problem, 0)
	for _, v := range p.data {
		if v.GetContestID() == id {
			res = append(res, v)
		}
	}
	return res, nil
}

func (p *ProblemRepository) CreateProblem(in domain.Problem) error {
	p.data = append(p.data, in)
	return nil
}

func (p *ProblemRepository) FindProblemByID(id id.SnowFlakeID) (*domain.Problem, error) {
	for _, v := range p.data {
		if v.GetProblemID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func (p *ProblemRepository) FindProblemByTitle(name string) (*domain.Problem, error) {
	for _, v := range p.data {
		if v.GetTitle() == name {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func (p *ProblemRepository) FindCaseSetByID(id id.SnowFlakeID) (*domain.Caseset, error) {
	for _, k := range p.data {
		for _, v := range k.GetCaseSets() {
			if v.GetID() == id {
				return &v, nil
			}
		}
	}
	return nil, errors.New("not found")
}

func (p *ProblemRepository) FindCaseByID(id id.SnowFlakeID) (*domain.Case, error) {
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
