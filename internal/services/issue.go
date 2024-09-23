package services

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"web-pet-project/internal/dbms/model"
	"web-pet-project/internal/dbms/repository"
)

type IssuesService struct {
	repos []repository.IssuesRepository
}

func NewIssuesService(rs []repository.IssuesRepository) IssuesService {
	return IssuesService{
		repos: rs,
	}
}

func (s *IssuesService) CloseRepos() {
	for _, s := range s.repos {
		s.Close()
	}
}

func (s *IssuesService) GetIssueListAsCsv() ([]byte, error) {

	allIssues := parallelGetIssues(s)

	var buf bytes.Buffer
	csvW := csv.NewWriter(&buf)
	csvW.Write(
		[]string{allIssues[0].IssueId, strconv.Itoa(allIssues[0].IssueType), allIssues[0].IssueKey, allIssues[0].Summary})
	for _, issue := range allIssues {
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

func (s *IssuesService) GetIssueListAsJson() ([]byte, error) {

	allIssues := parallelGetIssues(s)

	bs, err := json.Marshal(allIssues)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func parallelGetIssues(s *IssuesService) []model.Issue {
	var wg sync.WaitGroup
	var allIssues []model.Issue
	for _, v := range s.repos {
		wg.Add(1)
		v := v
		go func() {
			defer wg.Done()
			log.Printf("Get data from %v\n", v)
			issueList, err := v.GetAllIssues()
			if err != nil {
				log.Printf("Error getting data from :%v\n %v\n", v, err)
			} else {
				allIssues = append(allIssues, issueList...)
			}
		}()
	}
	wg.Wait()
	return allIssues
}
