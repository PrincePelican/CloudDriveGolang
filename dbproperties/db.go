package dbproperties

import (
	"database/sql"
	"fmt"
	"log"
)

func Connection() *sql.DB {
	host := "127.0.0.1"
	port := "5432"
	user := "postgres"
	password := "admin"
	dbname := "postgres"

	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psql)
	if err != nil {
		log.Fatalf("DataBase init error: %s", err)
	}

	return db
}
