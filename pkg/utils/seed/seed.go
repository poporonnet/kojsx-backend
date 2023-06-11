package seed

import (
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/domain"
	"github.com/mct-joken/kojs5-backend/pkg/utils/password/argon2"
)

type Seeds struct {
	Contests   []domain.Contest
	Users      []domain.User
	Problems   []domain.Problem
	Submission []domain.Submission
}

func NewSeeds() Seeds {
	loc := time.UTC
	d := time.Date(2021, time.October, 1, 0, 0, 0, 0, loc)
	passwordEncoder := argon2.NewArgon2PasswordEncoder()

	contest1 := domain.NewContest("10")
	_ = contest1.SetTitle("Contest 1")
	_ = contest1.SetDescription("# About\nThis Contest is for seed")
	_ = contest1.SetStartAt(d)
	_ = contest1.SetEndAt(d.Add(24 * time.Hour * 31 * 12 * 10))
	contests := []domain.Contest{*contest1}

	user1, _ := domain.NewUser("20", "Eric", "eric@example.jp")
	enc, _ := passwordEncoder.EncodePassword("294729dnr0@sn!")
	user1.SetPassword(string(enc))
	user1.SetVerified()
	user2, _ := domain.NewUser("30", "Eric", "eric@example.jp")
	enc2, _ := passwordEncoder.EncodePassword("294729dnr0@sn!")
	user2.SetPassword(string(enc2))
	user2.SetAdmin()
	users := []domain.User{*user1, *user2}

	problems := func() []domain.Problem {
		// CaseSets
		set1 := domain.NewCaseset("50")
		_ = set1.SetName("test1")
		_ = set1.SetPoint(200)
		set2 := domain.NewCaseset("60")
		_ = set2.SetName("test2")
		_ = set2.SetPoint(100)
		// Case
		case1 := domain.NewCase("70", set1.GetID())
		_ = case1.SetIn("hello\n")
		_ = case1.SetOut("world\n")
		case2 := domain.NewCase("80", set1.GetID())
		_ = case2.SetIn("1 2\n")
		_ = case2.SetOut("3\n")
		case3 := domain.NewCase("90", set2.GetID())
		_ = case3.SetIn("abc\n")
		_ = case3.SetOut("3\n")
		case4 := domain.NewCase("100", set2.GetID())
		_ = case4.SetIn("abc\nabp abc\n")
		_ = case4.SetOut("2\n")

		_ = set1.AddCase(*case1)
		_ = set1.AddCase(*case2)
		_ = set2.AddCase(*case3)
		_ = set2.AddCase(*case4)
		// Problems
		problem := domain.NewProblem("110", contest1.GetID())
		_ = problem.SetIndex("A")
		_ = problem.SetTitle("Moji")
		_ = problem.SetText("Calculate the number.\n")
		_ = problem.SetTimeLimit(2000)

		_ = problem.AddCaseSet(*set1)
		_ = problem.AddCaseSet(*set2)
		return []domain.Problem{*problem}
	}()

	submissions := func() []domain.Submission {
		submission1, _ := domain.NewSubmission(
			"200",
			problems[0].GetProblemID(),
			"1",
			"G++",
			"I2luY2x1ZGUgPGlvc3RyZWFtPgoKaW50IG1haW4oKSB7CiAgc3RkOjpjb3V0IDw8ICJ3b3JsZCIgPDwgc3RkOjplbmRsOwp9Cg==",
			d,
		)
		submission2, _ := domain.NewSubmission(
			"210",
			problems[0].GetProblemID(),
			"1",
			"G++",
			"I2luY2x1ZGUgPGlvc3RyZWFtPgoKaW50IG1haW4oKSB7CiAgc3RkOjpjb3V0IDw8ICJ3b3JsZCIgPDwgc3RkOjplbmRsOwp9Cg==",
			d,
		)
		return []domain.Submission{*submission1, *submission2}
	}()

	return Seeds{
		Contests:   contests,
		Users:      users,
		Problems:   problems,
		Submission: submissions,
	}
}
