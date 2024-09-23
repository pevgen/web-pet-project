package memory

import (
	"fmt"
	"web-pet-project/internal/dbms/model"
	"web-pet-project/internal/dbms/repository"
)

type issuesRepository struct {
}

func NewIssuesRepository() repository.IssuesRepository {
	return &issuesRepository{}
}

func (repo *issuesRepository) GetAllIssues() ([]model.Issue, error) {
	var issues = make([]model.Issue, 2)
	issues[0] = model.Issue{"1", "k1", 1, "Рус"}
	issues[1] = model.Issue{"2", "k2", 2, "Рус 2"}
	return issues, nil
}

func (repo *issuesRepository) Close() {
	fmt.Println("memory.IssueRepository closed")
}
