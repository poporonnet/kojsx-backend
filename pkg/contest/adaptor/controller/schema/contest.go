package schema

import "time"

type CreateContestRequestJSON struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
}

type CreateContestResponseJSON struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"startAt"`
	EndAt       time.Time `json:"endAt"`
}

type FindContestResponseJSON = CreateContestResponseJSON

type GetRankingResponseJSON struct {
	Rank    int
	Point   int
	User    RankingUser
	Results []RankingProblemResult
}

type RankingUser struct {
	ID   string
	Name string
}

type RankingProblemResult struct {
	ProblemID string
	Point     int
}
