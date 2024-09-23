package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"strconv"
	"web-pet-project/internal/dbms/model"
	"web-pet-project/internal/dbms/repository"
)

type IssuesRepository struct {
	db *sql.DB
}

func NewIssueRepository(uri string) repository.IssuesRepository {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err)
	}

	return &IssuesRepository{
		db: db,
	}
}

// urlExample := "postgres://username:password@localhost:5432/database_name"
//
//	urlDb := "postgresql://myuser:secret@localhost:5432/reportapp" // os.Getenv("DATABASE_URL")
//
// const (
//
//	host     = "localhost"
//	port     = 5432
//	user     = "myuser"
//	password = "secret"
//	dbname   = "reportapp"
//
// )
func (repo *IssuesRepository) Close() {
	repo.db.Close()
}

// var db *sql.DB
func (repo *IssuesRepository) GetAllIssues() ([]model.Issue, error) {

	err := repo.db.Ping()
	CheckError(err)

	fmt.Println("Successfully connected to Postgres!")

	count := 0
	rows, err := repo.db.Query("SELECT issue_id, issue_key,issue_type,summary FROM issues")
	CheckError(err)

	var result []model.Issue
	defer rows.Close()

	for rows.Next() {
		count++
		var issueId, issueKey, issueType, summary string
		err = rows.Scan(&issueId, &issueKey, &issueType, &summary)
		ii, _ := strconv.Atoi(issueType)
		result = append(result,
			model.Issue{IssueId: issueId, IssueKey: issueKey, IssueType: ii, Summary: summary})
		CheckError(err)
	}

	fmt.Printf("Load %d rows from DB", count)

	CheckError(err)

	return result, nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
