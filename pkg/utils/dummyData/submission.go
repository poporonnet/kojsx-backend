package dummyData

import (
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/domain"
)

var (
	NotExistsSubmission, _ = domain.NewSubmission("3", "2", "3", "Ruby", "p ARGV[2]", time.Now())
)
