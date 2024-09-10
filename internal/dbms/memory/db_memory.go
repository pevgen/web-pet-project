package memory

import "web-pet-project/internal/dbms"

type IssuesRepository struct {
}

func (mr *IssuesRepository) GetAllIssues() ([]dbms.Issue, error) {
	var issues = make([]dbms.Issue, 2)
	issues[0] = dbms.Issue{"1", "k1", "type 1", "Рус"}
	issues[1] = dbms.Issue{"2", "k2", "type 2", "Рус 2"}
	//issues = append(issues, Issue{"1", "k1", "type 1", "Рус"})
	//issues = append(issues, Issue{"2", "k2", "type 2", "Рус 2"})
	return issues, nil
}
