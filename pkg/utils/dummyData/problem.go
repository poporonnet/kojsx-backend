package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	e = domain.NewProblem("1", "1")
	_ = e.SetTitle("Test problem")
	f = domain.NewProblem("2", "1")
	_ = f.SetTitle("Test problem2")

	g                = domain.NewProblem("3", "1")
	_                = g.SetTitle("Test problem3")
	NotExistsProblem = g
	ExistsProblem    = e
	ProblemArray     = []domain.Problem{*e, *f}
)

