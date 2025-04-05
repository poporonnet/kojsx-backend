package dummyData

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
)

var (
	m = model.NewContestant("1", "2", "3")
	n = model.NewContestant("4", "2", "5")

	ExistsContestantData    = m
	NotExistsContestantData = model.NewContestant("6", "2", "7")

	ContestantArray = []model.Contestant{*m, *n}
)
