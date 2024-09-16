package repository

import "web-pet-project/internal/dbms/model"

type IssuesRepository interface {
	GetAllIssues() ([]model.Issue, error)
	//Save([]Issue) error
}
