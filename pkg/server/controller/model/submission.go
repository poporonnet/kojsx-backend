package model

import "time"

type CreateSubmissionRequestJSON struct {
	ProblemID string `json:"problemID"`
	Code      string `json:"code"`
	Lang      string `json:"lang"`
}

type CreateSubmissionResponseJSON struct {
	ID        string `json:"id"`
	ProblemID string `json:"problemID"`
	Code      string `json:"code"`
	Lang      string `json:"lang"`
}

type GetSubmissionTaskResponseJSON struct {
	ID        string                           `json:"ID"`
	ProblemID string                           `json:"problemID"`
	Lang      string                           `json:"lang"`
	Code      string                           `json:"Code"`
	Cases     []GetSubmissionTaskResponseCases `json:"cases"`
	Config    GetSubmissionTaskResponseConfig  `json:"config"`
}

type GetSubmissionTaskResponseCases struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type GetSubmissionTaskResponseConfig struct {
	TimeLimit   int `json:"timeLimit"`
	MemoryLimit int `json:"memoryLimit"`
}

type CreateSubmissionResultRequestJSON struct {
	SubmissionID        string                    `json:"submissionID"`
	ProblemID           string                    `json:"problemID"`
	LanguageType        string                    `json:"languageType"`
	CompilerMessage     string                    `json:"compilerMessage"`
	CompileErrorMessage string                    `json:"compileErrorMessage"`
	Results             []CreateSubmissionResults `json:"results"`
}

type CreateSubmissionResults struct {
	CaseName   string `json:"caseName"`
	Output     string `json:"output"`
	ExitStatus int    `json:"exitStatus"`
	Duration   int    `json:"duration"`
	Usage      int    `json:"usage"`
}

type GetSubmissionResponseJSON struct {
	ID          string    `json:"id"`
	SubmittedAt time.Time `json:"submittedAt"`
	User        struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	Problem struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"problem"`
	Code    string                 `json:"code"`
	Lang    string                 `json:"lang"`
	Points  int                    `json:"points"`
	Status  string                 `json:"status"`
	Time    int                    `json:"time"`
	Memory  int                    `json:"memory"`
	Results []GetSubmissionResults `json:"results"`
}

type GetSubmissionResults struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Time   int    `json:"time"`
	Memory int    `json:"memory"`
}

type FindSubmissionByContestIDResponseJSON struct {
	ID          string    `json:"id"`
	SubmittedAt time.Time `json:"submittedAt"`
	User        struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	Problem struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"problem"`
	Lang   string `json:"lang"`
	Points int    `json:"points"`
	Status string `json:"status"`
	Time   int    `json:"time"`
	Memory int    `json:"memory"`
}
