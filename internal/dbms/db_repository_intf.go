package dbms

type IssuesRepository interface {
	GetAllIssues() ([]Issue, error)
	Save([]Issue) error
}
