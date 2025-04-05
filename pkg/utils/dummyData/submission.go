package dummyData

import (
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/contest/model"
)

var (
	NotExistsSubmission, _ = model.NewSubmission("3", "2", "3", "Ruby", "p ARGV[2]", time.Now())
)
