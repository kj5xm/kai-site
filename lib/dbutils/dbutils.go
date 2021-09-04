package dbutil

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

var db *sql.DB

func GetDB() (*sql.DB, error) {
	//connectionString := "user=madhanganesh dbname=taskpad sslmode=disable"
	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connectionString == "" {
		return nil, errors.New("'POSTGRES_CONNECTION_STRING' environment variable not set")
	}

	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}
	return conn, nil
}
