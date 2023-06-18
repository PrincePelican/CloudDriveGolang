package dbproperties

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func Connection() *sql.DB {
	// local properties for database
	host := "localhost"
	port := "5432"
	dbname := "postgres"
	user := "postgres"
	password := "admin"

	initEnvVariables(&host, &port, &dbname, &user, &password)

	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psql)
	if err != nil {
		log.Fatalf("DataBase init error: %s", err)
	}

	return db
}

func initEnvVariables(host *string, port *string, dbname *string, user *string, password *string) {
	if env := os.Getenv("DbHostName"); env != "" {
		*host = env
	}
	if env := os.Getenv("DbPort"); env != "" {
		*port = env
	}
	if env := os.Getenv("DbName"); env != "" {
		*dbname = env
	}
	if env := os.Getenv("DbUser"); env != "" {
		*user = env
	}
	if env := os.Getenv("DbPasword"); env != "" {
		*password = env
	}
}
