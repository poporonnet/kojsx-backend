package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	g = domain.NewProblem("3", "1")
	_ = g.SetTitle("Test problem3")

	NotExistsProblem     = g
	NotExistsCasesetData = domain.NewCaseset("3")
	NotExistsCase        = domain.NewCase("5", "6")
)
