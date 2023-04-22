package submission

import (
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type Data struct {
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

	results []Result
}

type Result struct {
	result     string
	caseName   string
	execTime   int
	execMemory int
}

func newResult(resultString string, caseName string, execTime int, execMemory int) *Result {
	return &Result{result: resultString, caseName: caseName, execTime: execTime, execMemory: execMemory}
}

func (r Result) GetResult() string {
	return r.result
}

func (r Result) GetCaseName() string {
	return r.caseName
}

func (r Result) GetExecTime() int {
	return r.execTime
}

func (r Result) GetExecMemory() int {
	return r.execMemory
}

func NewData(
	id id.SnowFlakeID,
	problemID id.SnowFlakeID,
	contestantID id.SnowFlakeID,
	point int,
	lang string,
	codeLength int,
	result string,
	execTime int,
	execMemory int,
	code string,
	submittedAt time.Time,
	results []Result,
) *Data {
	return &Data{
		id:           id,
		problemID:    problemID,
		contestantID: contestantID,
		point:        point,
		lang:         lang,
		codeLength:   codeLength,
		result:       result,
		execTime:     execTime,
		execMemory:   execMemory,
		code:         code,
		submittedAt:  submittedAt,
		results:      results,
	}
}

func (d Data) GetID() id.SnowFlakeID {
	return d.id
}

func (d Data) GetProblemID() id.SnowFlakeID {
	return d.problemID
}

func (d Data) GetContestantID() id.SnowFlakeID {
	return d.contestantID
}

func (d Data) GetPoint() int {
	return d.point
}

func (d Data) GetLang() string {
	return d.lang
}

func (d Data) GetCodeLength() int {
	return d.codeLength
}

func (d Data) GetResult() string {
	return d.result
}

func (d Data) GetExecTime() int {
	return d.execTime
}

func (d Data) GetExecMemory() int {
	return d.execMemory
}

func (d Data) GetCode() string {
	return d.code
}

func (d Data) GetSubmittedAt() time.Time {
	return d.submittedAt
}

func (d Data) GetResults() []Result {
	return d.results
}

func DomainToData(in domain.Submission) *Data {
	return NewData(
		in.GetID(),
		in.GetProblemID(),
		in.GetContestantID(),
		in.GetPoint(),
		in.GetLang(),
		in.GetCodeLength(),
		in.GetResult(),
		in.GetExecTime(),
		in.GetExecMemory(),
		in.GetCode(),
		in.GetSubmittedAt(),
		submissionResultToResults(in.GetResults()),
	)
}

func submissionResultToResults(in []domain.SubmissionResult) []Result {
	res := make([]Result, len(in))
	for i, v := range in {
		res[i] = *newResult(v.GetResult(), v.GetCaseName(), v.GetExecTime(), v.GetExecMemory())
	}
	return res
}

func DataToDomain(in Data) *domain.Submission {
	r, _ := domain.NewSubmission(in.GetID(), in.GetProblemID(), in.GetContestantID(), in.GetLang(), in.GetCode(), in.GetSubmittedAt())
	addSubmissionResult(r, in.GetResults())
	return r
}

func resultToSubmissionResult(in []Result) []domain.SubmissionResult {
	res := make([]domain.SubmissionResult, len(in))
	return res
}

func addSubmissionResult(in *domain.Submission, results []Result) {
	r := resultToSubmissionResult(results)
	for i := range results {
		err := in.AddResult(r[i])
		if err != nil {
			return
		}
	}
}
