package dummyData

import (
	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
)

var (
	nc                   = model.NewContest("3")
	_                    = nc.SetTitle("Contest")
	NotExistsContestData = nc
)
