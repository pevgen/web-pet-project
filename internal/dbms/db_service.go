package dbms

import (
	"database/sql"
	"log"
	"os"
	"time"
)

func GetDBCredentials() (string, string) {
	dbType := os.Getenv("DB_TYPE")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	return dbType, dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
}

func SetDbConnectionPool() {
	db, err := sql.Open("postgres", "user=pqtest dbname=pqtest sslmode=verify-full")
	if err != nil {
		log.Fatal(err)
	}
	// Задаем максимальное количество открытых соединений
	db.SetMaxOpenConns(100)
	// Задаем максимальное количество простаивающих соединений
	db.SetMaxIdleConns(25)
	// Задаем время истечения соединения
	db.SetConnMaxLifetime(5 * time.Minute)
}
