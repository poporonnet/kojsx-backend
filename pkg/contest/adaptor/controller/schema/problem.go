package schema

type CreateProblemRequestJSON struct {
	ContestID string `json:"contestID"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Points    int    `json:"points"`
	Limits    struct {
		Memory int `json:"memory"`
		Time   int `json:"time"`
	} `json:"limits"`
}

type CreateProblemResponseJSON struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Points int    `json:"points"`
	Limits struct {
		Memory int `json:"memory"`
		Time   int `json:"time"`
	} `json:"limits"`
}

type FindProblemResponseJSON = CreateProblemResponseJSON
