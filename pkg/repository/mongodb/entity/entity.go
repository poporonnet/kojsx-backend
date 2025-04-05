package entity

import (
	"time"

	"github.com/poporonnet/kojsx-backend/pkg/domain"
	"github.com/poporonnet/kojsx-backend/pkg/utils/id"
)

type CaseSet struct {
	ID    id.SnowFlakeID
	Name  string
	Point int

	Cases []Case `bson:"cases"`
}

func (c CaseSet) ToDomain() domain.Caseset {
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

func (c Case) ToDomain() domain.Case {
	ca := domain.NewCase(c.ID, c.CaseSetID)
	_ = ca.SetIn(c.In)
	_ = ca.SetOut(c.Out)

	return *ca
}

type Problem struct {
	ID        id.SnowFlakeID `bson:"_id"`
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
	pr := domain.NewProblem(p.ID, p.ContestID)
	_ = pr.SetTitle(p.Title)
	_ = pr.SetIndex(p.Index)
	_ = pr.SetText(p.Text)
	_ = pr.SetTimeLimit(p.TimeLimit)

	for _, v := range p.CaseSets {
		sets := v.ToDomain()
		for _, k := range v.Cases {
			_ = sets.AddCase(k.ToDomain())
		}
		_ = pr.AddCaseSet(sets)
	}

	return *pr
}

type Contestant struct {
	ID        id.SnowFlakeID `bson:"_id"`
	ContestID id.SnowFlakeID `bson:"contestID"`
	UserID    id.SnowFlakeID `bson:"userID"`

	Role  int
	Point int
}

func (c Contestant) ToDomain() domain.Contestant {
	co := domain.NewContestant(c.ID, c.ContestID, c.UserID)
	if c.Role == 1 {
		co.SetAdmin()
	}
	_ = co.SetPoint(c.Point)

	return *co
}

type Contest struct {
	ID id.SnowFlakeID `bson:"_id"`

	Title       string
	Description string
	StartAt     time.Time `bson:"startAt"`
	EndAt       time.Time `bson:"endAt"`
}

func (c Contest) ToDomain() domain.Contest {
	co := domain.NewContest(c.ID)
	_ = co.SetTitle(c.Title)
	_ = co.SetStartAt(c.StartAt)
	_ = co.SetEndAt(c.StartAt)
	return *co
}

type User struct {
	ID       id.SnowFlakeID `bson:"_id"`
	Name     string
	Email    string
	Password string
	Role     int
}

func (u User) ToDomain() domain.User {
	ur, _ := domain.NewUser(u.ID, u.Name, u.Email)
	switch u.Role {
	case 0:
		ur.SetAdmin()
	case 1:
		ur.SetNormal()
	}
	ur.SetPassword(u.Password)
	return *ur
}

type Submission struct {
	ID           id.SnowFlakeID `bson:"_id"`
	ProblemID    id.SnowFlakeID `bson:"problemID"`
	ContestantID id.SnowFlakeID `bson:"contestantID"`

	Point       int
	Lang        string
	CodeLength  int `bson:"codeLength"`
	Result      string
	ExecTime    int `bson:"execTime"`
	ExecMemory  int `bson:"execMemory"`
	Code        string
	SubmittedAt time.Time `bson:"submittedAt"`

	Results []SubmissionResult
}

func (s Submission) ToDomain() domain.Submission {
	sb, _ := domain.NewSubmission(s.ID, s.ProblemID, s.ContestantID, s.Lang, s.Code, s.SubmittedAt)
	_ = sb.SetPoint(s.Point)
	sb.SetExecMemory(s.ExecMemory)
	sb.SetExecTime(s.ExecTime)
	sb.SetResult(s.Result)

	for _, v := range s.Results {
		_ = sb.AddResult(v.toDomain())
	}
	return *sb
}

type SubmissionResult struct {
	ID         id.SnowFlakeID
	Result     string
	Output     string
	CaseName   string `bson:"caseName"`
	ExitStatus int    `bson:"exitStatus"`
	ExecTime   int    `bson:"execTime"`
	ExecMemory int    `bson:"execMemory"`
}

func (s SubmissionResult) toDomain() domain.SubmissionResult {
	return *domain.NewSubmissionResult(
		s.ID,
		s.Result,
		s.Output,
		s.CaseName,
		s.ExitStatus,
		s.ExecTime,
		s.ExecMemory,
	)
}
