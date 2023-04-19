package dummyData

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"time"
)

var (
	o, _ = domain.NewSubmission("1", "2", "3", "Ruby", "p ARGV[0]", time.Now())
	p, _ = domain.NewSubmission("2", "2", "3", "Ruby", "p ARGV[1]", time.Now())

	NotExistsSubmission, _ = domain.NewSubmission("3", "2", "3", "Ruby", "p ARGV[2]", time.Now())
	ExistsSubmission       = o

	SubmissionArray = []domain.Submission{*o, *p}
)

