package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	nc                   = domain.NewContest("3")
	_                    = nc.SetTitle("Contest")
	NotExistsContestData = nc
)
