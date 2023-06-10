package dummyData

import "github.com/mct-joken/kojs5-backend/pkg/domain"

var (
	q = domain.NewSubmissionResult("1", "AC", "hello world\n", "test1.txt", 0, 10, 1000)
	r = domain.NewSubmissionResult("2", "AC", "", "test2.txt", 0, 20, 2000)

	NotExistsSubmissionResult = domain.NewSubmissionResult("3", "AC", "", "test3.txt", -1, 30, 2500)
	ExistsSubmissionResult    = q

	SubmissionResultArray = []domain.SubmissionResult{*q, *r}
)
