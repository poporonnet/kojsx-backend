package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	h = domain.NewCase("1", "2")
	j = domain.NewCase("3", "4")

	NotExistsCase = domain.NewCase("5", "6")
	ExistsCase    = h

	CaseArray = []domain.Case{*h, *j}
)

