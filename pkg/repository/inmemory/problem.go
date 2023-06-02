package inmemory

import (
	"errors"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ProblemRepository struct {
	data     []domain.Problem
	sets     []domain.Caseset
	caseData []domain.Case
}

func NewProblemRepository(data []domain.Problem, sets []domain.Caseset, caseData []domain.Case) *ProblemRepository {
	return &ProblemRepository{data: data, sets: sets, caseData: caseData}
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
	for _, v := range p.sets {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}

func (p *ProblemRepository) FindCaseByID(id id.SnowFlakeID) (*domain.Case, error) {
	for _, v := range p.caseData {
		if v.GetID() == id {
			return &v, nil
		}
	}
	return nil, errors.New("not found")
}
