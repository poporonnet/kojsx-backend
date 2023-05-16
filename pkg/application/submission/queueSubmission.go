package submission

import (
	"errors"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"sort"
	"time"
)

/*
memo: CreateServiceから呼ばれるサービスなので基本的に公開しない

*/

// QueueSubmissionService 提出の実行/実行結果関連のサービス
type QueueSubmissionService struct {
	tasks []RunnerTask
}

func NewQueueSubmissionService() *QueueSubmissionService {
	t := make([]RunnerTask, 100)
	for i := range t {
		t[i].IsWaiting = false
		t[i].IsJobFinished = true
	}
	return &QueueSubmissionService{tasks: t}
}

type Language string

// RunnerTask SubmissionRunnerTask agentに渡す実行するためのデータ
type RunnerTask struct {
	SubmissionID id.SnowFlakeID    `json:"submissionID"`
	ProblemID    id.SnowFlakeID    `json:"problemID"`
	Lang         Language          `json:"lang"`
	Code         string            `json:"code"`
	Cases        []submissionCases `json:"cases"`
	TimeLimit    int               `json:"timeLimit"`
	MemoryLimit  int               `json:"memoryLimit"`

	SubmittedAt   time.Time
	IsJobFinished bool
	IsWaiting     bool
}

type submissionCases struct {
	Name string
	File string
}

// AddTask タスクを追加
func (s *QueueSubmissionService) AddTask(t RunnerTask) error {
	for i, v := range s.tasks {
		if v.IsJobFinished {
			s.tasks[i] = t
			return nil
		}
	}
	return errors.New("failed to add task")
}

func (s *QueueSubmissionService) PopTask() (*RunnerTask, error) {
	tasks := s.tasks
	// First In First Outなので時系列順にソート
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].SubmittedAt.Before(tasks[j].SubmittedAt)
	})

	for i, v := range tasks {
		// Jobが終了しておらず、待機中になっているものをとる
		if !v.IsJobFinished && v.IsWaiting {
			res := tasks[i]

			// 実行中にする
			tasks[i].IsWaiting = false
			return &res, nil
		}
	}

	return nil, errors.New("nothing to pop")
}

func (s *QueueSubmissionService) GetQueueTasks() *[]RunnerTask {
	return &s.tasks
}
