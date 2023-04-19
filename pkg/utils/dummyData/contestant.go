package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	m = domain.NewContestant("1", "2", "3")
	n = domain.NewContestant("4", "2", "5")

	ExistsContestantData    = m
	NotExistsContestantData = domain.NewContestant("6", "2", "7")

	ContestantArray = []domain.Contestant{*m, *n}
)

