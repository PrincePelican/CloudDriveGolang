package dbproperties

import (
	"cloud-service/entity"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	// local properties for database
	host := "localhost"
	port := "5432"
	dbname := "postgres"
	user := "postgres"
	password := "admin"

	initEnvVariables(&host, &port, &dbname, &user, &password)

	psql := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psql), &gorm.Config{})
	if err != nil {
		log.Fatalf("DataBase init error: %s", err)
	}

	return db
}

func InitTables(db *gorm.DB) {
	db.AutoMigrate(&entity.ResourceEntity{})
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
