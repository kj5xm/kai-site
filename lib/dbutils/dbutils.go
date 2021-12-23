package dbutil

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "testdb"
)

var db *sql.DB

func GetDB() (*sql.DB, error) {
	//connectionString := "user=madhanganesh dbname=taskpad sslmode=disable"
	//connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	//if connectionString == "" {
	//	return nil, errors.New("'POSTGRES_CONNECTION_STRING' environment variable not set")
	//}
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(fmt.Sprintf("DB: %v", err))
	}
	return conn, nil
}
