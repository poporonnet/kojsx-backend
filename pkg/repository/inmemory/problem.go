package inmemory

import (
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

func (p *ProblemRepository) CreateProblem(in domain.Problem) error {
	p.data = append(p.data, in)
	return nil
}

func (p *ProblemRepository) FindProblemByID(id id.SnowFlakeID) *domain.Problem {
	for _, v := range p.data {
		if v.GetProblemID() == id {
			return &v
		}
	}
	return nil
}

func (p *ProblemRepository) FindProblemByTitle(name string) *domain.Problem {
	for _, v := range p.data {
		if v.GetTitle() == name {
			return &v
		}
	}
	return nil
}

func (p *ProblemRepository) FindCaseSetByID(id id.SnowFlakeID) *domain.Caseset {
	for _, v := range p.sets {
		if v.GetID() == id {
			return &v
		}
	}
	return nil
}

func (p *ProblemRepository) FindCaseByID(id id.SnowFlakeID) *domain.Case {
	for _, v := range p.caseData {
		if v.GetID() == id {
			return &v
		}
	}
	return nil
}
