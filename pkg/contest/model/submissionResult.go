package model

import "github.com/poporonnet/kojsx-backend/pkg/utils/id"

// SubmissionResult 提出詳細(ケースごとの結果/ケース名/実行時間/実行メモリ
type SubmissionResult struct {
	id         id.SnowFlakeID
	output     string
	result     string
	caseName   string
	exitStatus int
	execTime   int
	execMemory int
}

func (s *SubmissionResult) GetOutput() string {
	return s.output
}

func (s *SubmissionResult) GetExitStatus() int {
	return s.exitStatus
}

func (s *SubmissionResult) GetID() id.SnowFlakeID {
	return s.id
}

func (s *SubmissionResult) GetResult() string {
	return s.result
}

func (s *SubmissionResult) GetCaseName() string {
	return s.caseName
}

func (s *SubmissionResult) GetExecTime() int {
	return s.execTime
}

func (s *SubmissionResult) GetExecMemory() int {
	return s.execMemory
}

// NewSubmissionResult 不変値: すべて
func NewSubmissionResult(id id.SnowFlakeID, result, output, caseName string, exitStatus, execTime, execMemory int) *SubmissionResult {
	return &SubmissionResult{
		id:         id,
		output:     output,
		result:     result,
		caseName:   caseName,
		exitStatus: exitStatus,
		execTime:   execTime,
		execMemory: execMemory,
	}
}
