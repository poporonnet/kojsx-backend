package entity

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"time"
)

type CaseSet struct {
	ID    id.SnowFlakeID
	Name  string
	Point int

	Cases []Case
}

func (c CaseSet) toDomain() domain.Caseset {
	cs := domain.NewCaseset(c.ID)
	_ = cs.SetName(c.Name)
	_ = cs.SetPoint(c.Point)

	return *cs
}

type Case struct {
	ID        id.SnowFlakeID
	CaseSetID id.SnowFlakeID
	In        string
	Out       string
}

func (c Case) toDomain() domain.Case {
	ca := domain.NewCase(c.ID, c.CaseSetID)
	_ = ca.SetIn(c.In)
	_ = ca.SetOut(c.Out)

	return *ca
}

type Problem struct {
	ID        id.SnowFlakeID
	ContestID id.SnowFlakeID `bson:"contestID"`

	Index       string
	Title       string
	Text        string
	Point       int
	MemoryLimit int `bson:"memoryLimit"`
	TimeLimit   int `bson:"timeLimit"`

	CaseSets []CaseSet
}

func (p Problem) ToDomain() domain.Problem {

}

type Contestant struct {
	ID        id.SnowFlakeID
	ContestID id.SnowFlakeID `bson:"contestID"`
	UserID    id.SnowFlakeID `bson:"userID"`

	Role  int
	Point int
}

type Contest struct {
	ID id.SnowFlakeID

	Title       string
	Description string
	StartAt     string `bson:"startAt"`
	EndAt       string `bson:"endAt"`
}

type User struct {
	ID       id.SnowFlakeID
	Name     string
	Email    string
	Password string
	Role     int
}

type Submission struct {
	ID           id.SnowFlakeID
	ProblemID    id.SnowFlakeID
	ContestantID id.SnowFlakeID `bson:"contestantID"`

	Point      int
	Lang       string
	CodeLength int
	Result     string
	ExecTime   int `bson:"execTime"`
	ExecMemory int `bson:"execMemory"`
	Code       string

	SubmittedAt time.Time

	Results []SubmissionResult
}

type SubmissionResult struct {
	ID           id.SnowFlakeID
	SubmissionID id.SnowFlakeID `bson:"submissionID"`
	Result       string
	CaseName     string `bson:"caseName"`
	ExecTime     int    `bson:"execTime"`
	ExecMemory   int    `bson:"execMemory"`
}
