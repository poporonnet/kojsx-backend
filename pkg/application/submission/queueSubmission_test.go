package submission

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var s *QueueSubmissionService

func init() {
	s = NewQueueSubmissionService()
}

func TestQueueSubmissionService_AddTask(t *testing.T) {
	err := s.AddTask(RunnerTask{
		SubmissionID:  "123",
		ProblemID:     "123",
		Lang:          "G++",
		Code:          "I2luY2x1ZGUgPGlvc3RyZWFtPgoKaW50IG1haW4oKSB7CglzdGQ6OmNvdXQgPDwgIkhlbGxvIFdvcmxkISIgPDwgc3RkOjplbmRsOwp9Cgo=",
		Cases:         nil,
		TimeLimit:     0,
		MemoryLimit:   0,
		SubmittedAt:   time.Time{},
		IsJobFinished: false,
		IsWaiting:     true,
	})
	assert.Equal(t, nil, err)
}

func TestQueueSubmissionService_PopTask(t *testing.T) {
	_ = s.AddTask(RunnerTask{
		SubmissionID:  "123",
		ProblemID:     "123",
		Lang:          "G++",
		Code:          "I2luY2x1ZGUgPGlvc3RyZWFtPgoKaW50IG1haW4oKSB7CglzdGQ6OmNvdXQgPDwgIkhlbGxvIFdvcmxkISIgPDwgc3RkOjplbmRsOwp9Cgo=",
		Cases:         nil,
		TimeLimit:     0,
		MemoryLimit:   0,
		SubmittedAt:   time.Time{},
		IsJobFinished: false,
		IsWaiting:     true,
	})

	task, err := s.PopTask()
	if err != nil {
		return
	}

	assert.Equal(t, &RunnerTask{
		SubmissionID:  "123",
		ProblemID:     "123",
		Lang:          "G++",
		Code:          "I2luY2x1ZGUgPGlvc3RyZWFtPgoKaW50IG1haW4oKSB7CglzdGQ6OmNvdXQgPDwgIkhlbGxvIFdvcmxkISIgPDwgc3RkOjplbmRsOwp9Cgo=",
		Cases:         nil,
		TimeLimit:     0,
		MemoryLimit:   0,
		SubmittedAt:   time.Time{},
		IsJobFinished: false,
		IsWaiting:     true,
	}, task)
}
