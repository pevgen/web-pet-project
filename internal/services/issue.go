package services

import (
	"bytes"
	"encoding/csv"
	"log"
)

type Issue struct {
	IssueId   string `json:"issueId"`
	IssueKey  string `json:"issueKey"`
	IssueType string `json:"issueType"`
	Summary   string `json:"summary"`
}

func GetIssueList() []Issue {
	var issues = make([]Issue, 2)
	issues[0] = Issue{"1", "k1", "type 1", "Рус"}
	issues[1] = Issue{"2", "k2", "type 2", "Рус 2"}
	//issues = append(issues, Issue{"1", "k1", "type 1", "Рус"})
	//issues = append(issues, Issue{"2", "k2", "type 2", "Рус 2"})
	return issues
}

func GetIssueListAsCsvBytes() []byte {
	var buf bytes.Buffer

	csvW := csv.NewWriter(&buf)

	issueList := GetIssueList()
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

	return buf.Bytes()
}
