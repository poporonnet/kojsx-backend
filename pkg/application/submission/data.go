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
}

type result struct {
	result     string
	caseName   string
	execTime   int
	execMemory int
}

func newResult(resultString string, caseName string, execTime int, execMemory int) *result {
	return &result{result: resultString, caseName: caseName, execTime: execTime, execMemory: execMemory}
}

func (r result) GetResult() string {
	return r.result
}

func (r result) GetCaseName() string {
	return r.caseName
}

func (r result) GetExecTime() int {
	return r.execTime
}

func (r result) GetExecMemory() int {
	return r.execMemory
}

func NewData(id id.SnowFlakeID, problemID id.SnowFlakeID, contestantID id.SnowFlakeID, point int, lang string, codeLength int, result string, execTime int, execMemory int, code string, submittedAt time.Time) *Data {
	return &Data{id: id, problemID: problemID, contestantID: contestantID, point: point, lang: lang, codeLength: codeLength, result: result, execTime: execTime, execMemory: execMemory, code: code, submittedAt: submittedAt}
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

func DomainToData(submission domain.Submission, result []domain.SubmissionResult) {
	
}
