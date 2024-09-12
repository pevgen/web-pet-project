package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"web-pet-project/internal/dbms"
)

type IssuesRepository struct {
	psqlInfo string
}

func NewIssueRepository() dbms.IssuesRepository {
	return &IssuesRepository{
		psqlInfo: fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname),
	}
}

// urlExample := "postgres://username:password@localhost:5432/database_name"
//
//	urlDb := "postgresql://myuser:secret@localhost:5432/reportapp" // os.Getenv("DATABASE_URL")
const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "secret"
	dbname   = "reportapp"
)

// var db *sql.DB
func (repo *IssuesRepository) GetAllIssues() ([]dbms.Issue, error) {
	db, err := sql.Open("postgres", repo.psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Successfully connected!")

	count := 0
	rows, err := db.Query("SELECT issue_id, issue_key,issue_type,summary FROM issues")
	CheckError(err)

	result := []dbms.Issue{}
	defer rows.Close()
	for rows.Next() {
		count++
		var issueId, issueKey, issueType, summary string
		err = rows.Scan(&issueId, &issueKey, &issueType, &summary)
		result = append(result,
			dbms.Issue{IssueId: issueId, IssueKey: issueKey, IssueType: issueType, Summary: summary})
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
