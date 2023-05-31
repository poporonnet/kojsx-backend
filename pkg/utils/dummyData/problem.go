package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	h = domain.NewCase("1", "2")
	j = domain.NewCase("3", "4")

	k = domain.NewCaseset("1")
	_ = k.AddCase(*h)
	l = domain.NewCaseset("2")

	e = domain.NewProblem("1", "1")
	_ = e.SetTitle("Test problem")
	_ = e.AddCaseSet(*k)
	f = domain.NewProblem("2", "1")
	_ = f.SetTitle("Test problem2")

	CaseArray    = []domain.Case{*h, *j}
	CasesetArray = []domain.Caseset{*k, *l}
	ProblemArray = []domain.Problem{*e, *f}

	g = domain.NewProblem("3", "1")
	_ = g.SetTitle("Test problem3")

	NotExistsProblem     = g
	ExistsProblem        = e
	ExistsCasesetData    = k
	NotExistsCasesetData = domain.NewCaseset("3")
	NotExistsCase        = domain.NewCase("5", "6")
	ExistsCase           = h
)
