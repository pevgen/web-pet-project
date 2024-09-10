package services

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"

	"web-pet-project/internal/dbms/postgres"
)

// var repo memory.IssuesRepository
var repo postgres.IssuesRepository

func GetIssueListAsCsv() ([]byte, error) {

	var buf bytes.Buffer

	csvW := csv.NewWriter(&buf)

	issueList, err := repo.GetAllIssues()
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

func GetIssueListAsJson() ([]byte, error) {

	issueList, _ := repo.GetAllIssues()
	bytes, err := json.Marshal(issueList)
	if err != nil {
		return nil, err

	}
	return bytes, nil
}
