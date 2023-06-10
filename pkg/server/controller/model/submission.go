package model

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
