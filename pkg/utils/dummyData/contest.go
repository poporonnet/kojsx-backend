package dummyData

import "github.com/poporonnet/kojsx-backend/pkg/domain"

var (
	nc                   = domain.NewContest("3")
	_                    = nc.SetTitle("Contest")
	NotExistsContestData = nc
)
