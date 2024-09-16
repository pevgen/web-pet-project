package model

type Issue struct {
	IssueId   string `json:"issueId"`
	IssueKey  string `json:"issueKey"`
	IssueType int    `json:"issueType"`
	Summary   string `json:"summary"`
}
