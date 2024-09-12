package services

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"
	"web-pet-project/internal/dbms"
)

type IssuesService struct {
	repo dbms.IssuesRepository
}

func NewIssuesService(repo dbms.IssuesRepository) IssuesService {
	return IssuesService{
		repo: repo,
	}
}

func (service *IssuesService) GetIssueListAsCsv() ([]byte, error) {

	var buf bytes.Buffer

	csvW := csv.NewWriter(&buf)

	issueList, err := service.repo.GetAllIssues()
	if err != nil {
		log.Fatal(err)
	}

	csvW.Write(
		[]string{issueList[0].IssueId, issueList[0].IssueType, issueList[0].IssueKey, issueList[0].Summary})
	for _, issue := range issueList {
		if err := csvW.Write(
			[]string{issue.IssueId, issue.IssueType, issue.IssueKey, issue.Summary}); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	csvW.Flush()
	if err := csvW.Error(); err != nil {
		log.Fatal(err)
	}

	return buf.Bytes(), nil
}

func (service *IssuesService) GetIssueListAsJson() ([]byte, error) {

	issueList, _ := service.repo.GetAllIssues()
	bytes, err := json.Marshal(issueList)
	if err != nil {
		return nil, err

	}
	return bytes, nil
}
