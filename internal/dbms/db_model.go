package dbms

type Issue struct {
	IssueId   string `json:"issueId"`
	IssueKey  string `json:"issueKey"`
	IssueType string `json:"issueType"`
	Summary   string `json:"summary"`
}
