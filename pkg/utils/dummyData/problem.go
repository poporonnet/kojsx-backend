package dummyData

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
)

var (
	g = model.NewProblem("3", "1")
	_ = g.SetTitle("Test problem3")

	NotExistsProblem     = g
	NotExistsCasesetData = model.NewCaseset("3")
	NotExistsCase        = model.NewCase("5", "6")
)
