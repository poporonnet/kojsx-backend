package entity

import (
	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
	"time"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
	Role     int
}

func (u User) ToDomain() domain.User {
	user, _ := domain.NewUser(id.SnowFlakeID(u.ID), u.Name, u.Email)
	if u.Role == domain.Normal {
		user.SetVerified()
	}
	if u.Role == domain.Admin {
		user.SetAdmin()
	}
	user.SetPassword(u.Password)

	return *user
}

type Contest struct {
	ID          string
	Title       string
	Description string
	StartAt     time.Time
	EndAt       time.Time
}

func (c Contest) ToDomain() domain.Contest {
	contest := domain.NewContest(id.SnowFlakeID(c.ID))
	_ = contest.SetTitle(c.Title)
	_ = contest.SetDescription(c.Description)
	_ = contest.SetStartAt(c.StartAt)
	_ = contest.SetEndAt(c.EndAt)
	return *contest
}

type Contestant struct {
	ID    string
	Role  int
	Point int

	ContestID string
	UserID    string
}

func (c Contestant) ToDomain() domain.Contestant {
	contestant := domain.NewContestant(id.SnowFlakeID(c.ID), id.SnowFlakeID(c.ContestID), id.SnowFlakeID(c.UserID))
	_ = contestant.SetPoint(c.Point)
	switch c.Role {
	case domain.ContestAdmin:
		contestant.SetAdmin()
	case domain.ContestParticipants:
		contestant.SetNormal()
	case domain.ContestTester:
		contestant.SetTester()
	}
	return *contestant
}

type Problem struct {
	ID          string
	Index       string
	Title       string
	Text        string
	Point       int
	MemoryLimit int
	TimeLimit   int

	ContestID string

	CaseSets []CaseSet
}

func (p Problem) ToDomain() domain.Problem {
	problem := domain.NewProblem(
		id.SnowFlakeID(p.ID),
		id.SnowFlakeID(p.ContestID),
	)
	_ = problem.SetIndex(p.Index)
	_ = problem.SetTitle(p.Title)
	_ = problem.SetText(p.Text)
	for _, v := range p.CaseSets {
		_ = v.ToDomain()
	}
	return domain.Problem{}
}

type Case struct {
	ID        string
	Input     string
	Output    string
	CasesetID string
}

func (c Case) ToDomain() domain.Case {
	d := domain.NewCase(id.SnowFlakeID(c.ID), id.SnowFlakeID(c.CasesetID))
	_ = d.SetIn(c.Input)
	_ = d.SetOut(c.Output)
	return *d
}

type CaseSet struct {
	ID        string
	Name      string
	Point     int
	ProblemID string

	Cases []Case
}

func (s CaseSet) ToDomain() domain.Caseset {
	set := domain.NewCaseset(id.SnowFlakeID(s.ID))
	_ = set.SetName(s.Name)
	_ = set.SetPoint(s.Point)
	for _, v := range s.Cases {
		_ = set.AddCase(v.ToDomain())
	}
	return *set
}

type Submission struct {
	ID          string
	Point       int
	Lang        string
	CodeLength  int
	Result      string
	ExecTime    int
	ExecMemory  int
	Code        string
	SubmittedAt time.Time

	ProblemID    string
	ContestantID string

	Results []SubmissionResult
}

func (s Submission) ToDomain() domain.Submission {
	sub, _ := domain.NewSubmission(id.SnowFlakeID(s.ID), id.SnowFlakeID(s.ProblemID), id.SnowFlakeID(s.ContestantID), s.Lang, s.Code, s.SubmittedAt)
	_ = sub.SetPoint(s.Point)
	sub.SetExecTime(s.ExecTime)
	sub.SetExecMemory(s.ExecMemory)
	sub.SetResult(s.Result)
	for _, v := range s.Results {
		_ = sub.AddResult(v.ToDomain())
	}

	return *sub
}

type SubmissionResult struct {
	ID         string
	Result     string
	Output     string
	CaseName   string
	ExitStatus int
	ExecTime   int
	ExecMemory int

	SubmissionID string
}

func (s SubmissionResult) ToDomain() domain.SubmissionResult {
	result := domain.NewSubmissionResult(id.SnowFlakeID(s.ID), s.Result, s.Output, s.CaseName, s.ExitStatus, s.ExecTime, s.ExecMemory)

	return *result
}
