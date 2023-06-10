package domain

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type Submission struct {
	id           id.SnowFlakeID
	problemID    id.SnowFlakeID
	contestantID id.SnowFlakeID
	point        int
	lang         string
	codeLength   int
	result       string
	execTime     int
	execMemory   int
	code         string
	submittedAt  time.Time

	results []SubmissionResult
}

func (s *Submission) GetResults() []SubmissionResult {
	return s.results
}

func (s *Submission) AddResult(result SubmissionResult) error {
	for _, v := range s.results {
		if v.GetID() == result.GetID() {
			return errors.New("AlreadyAdded")
		}
	}
	result.setSubmissionID(s.id)
	s.results = append(s.results, result)
	return nil
}

func (s *Submission) GetID() id.SnowFlakeID {
	return s.id
}

func (s *Submission) GetProblemID() id.SnowFlakeID {
	return s.problemID
}

func (s *Submission) GetContestantID() id.SnowFlakeID {
	return s.contestantID
}

func (s *Submission) GetPoint() int {
	return s.point
}

func (s *Submission) GetLang() string {
	return s.lang
}

func (s *Submission) GetCodeLength() int {
	return s.codeLength
}

func (s *Submission) GetResult() string {
	return s.result
}

func (s *Submission) GetExecTime() int {
	return s.execTime
}

func (s *Submission) GetExecMemory() int {
	return s.execMemory
}

func (s *Submission) GetCode() string {
	return s.code
}

func (s *Submission) GetSubmittedAt() time.Time {
	return s.submittedAt
}

func (s *Submission) SetPoint(point int) error {
	// 0~5000点, 100点刻み
	if point < 0 || point > 5000 || point%100 != 0 {
		return errors.New("InvalidPoint")
	}
	s.point = point
	return nil
}

func (s *Submission) SetResult(result string) {
	s.result = result
}

func (s *Submission) SetExecTime(execTime int) {
	s.execTime = execTime
}

func (s *Submission) SetExecMemory(execMemory int) {
	s.execMemory = execMemory
}

/*
NewSubmission
不変値: ID/ProblemID/ContestantID/Lang/CodeLength/Code/SubmittedAt
*/
func NewSubmission(id id.SnowFlakeID, pID id.SnowFlakeID, cID id.SnowFlakeID, lang string, code string, submittedAt time.Time) (*Submission, error) {
	return &Submission{
		id:           id,
		problemID:    pID,
		contestantID: cID,
		point:        0,
		lang:         lang,
		codeLength:   utf8.RuneCountInString(code),
		result:       "WE",
		execTime:     0,
		execMemory:   0,
		code:         code,
		submittedAt:  submittedAt,
	}, nil
}
