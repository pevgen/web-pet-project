package rdbms

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

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

func WorkWithDb() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT title, artist FROM album WHERE title = $1", "Jeru")
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var title string
		var artist string

		err = rows.Scan(&title, &artist)
		CheckError(err)

		fmt.Println(title, artist)
	}

	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
