package domain

import "github.com/mct-joken/kojs5-backend/pkg/utils/id"

// SubmissionResult 提出詳細(ケースごとの結果/ケース名/実行時間/実行メモリ
type SubmissionResult struct {
	id           id.SnowFlakeID
	submissionID id.SnowFlakeID
	result       string
	caseName     string
	execTime     int
	execMemory   int
}

func (s SubmissionResult) GetID() id.SnowFlakeID {
	return s.id
}

func (s SubmissionResult) GetSubmissionID() id.SnowFlakeID {
	return s.submissionID
}

func (s SubmissionResult) GetResult() string {
	return s.result
}

func (s SubmissionResult) GetCaseName() string {
	return s.caseName
}

func (s SubmissionResult) GetExecTime() int {
	return s.execTime
}

func (s SubmissionResult) GetExecMemory() int {
	return s.execMemory
}

// NewSubmissionResult 不変値: すべて
func NewSubmissionResult(id id.SnowFlakeID, submissionID id.SnowFlakeID, result string, caseName string, execTime int, execMemory int) *SubmissionResult {
	return &SubmissionResult{
		id:           id,
		submissionID: submissionID,
		result:       result,
		caseName:     caseName,
		execTime:     execTime,
		execMemory:   execMemory,
	}
}
