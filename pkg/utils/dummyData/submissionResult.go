package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	q = domain.NewSubmissionResult("1", "2", "AC", "test1.txt", 10, 1000)
	r = domain.NewSubmissionResult("2", "3", "AC", "test2.txt", 20, 2000)

	NotExistsSubmissionResult = domain.NewSubmissionResult("3", "4", "AC", "test3.txt", 30, 2500)
	ExistsSubmissionResult    = q

	SubmissionResultArray = []domain.SubmissionResult{*q, *r}
)
