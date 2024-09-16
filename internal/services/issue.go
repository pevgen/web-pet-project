package services

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"
	"strconv"
	"web-pet-project/internal/dbms/repository"
)

type IssuesService struct {
	repo repository.IssuesRepository
}

func NewIssuesService(repo repository.IssuesRepository) IssuesService {
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
		[]string{issueList[0].IssueId, strconv.Itoa(issueList[0].IssueType), issueList[0].IssueKey, issueList[0].Summary})
	for _, issue := range issueList {
		if err := csvW.Write(
			[]string{issue.IssueId, strconv.Itoa(issue.IssueType), issue.IssueKey, issue.Summary}); err != nil {
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
