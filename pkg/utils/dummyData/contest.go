package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	c = domain.NewContest("1")
	d = domain.NewContest("2")
	_ = c.SetTitle("Hello")
	_ = d.SetTitle("World")

	nc                   = domain.NewContest("3")
	_                    = nc.SetTitle("Contest")
	NotExistsContestData = nc
	ExistsContestData    = c

	ContestArray = []domain.Contest{*c, *d}
)
