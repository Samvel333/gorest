package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/samvel333/gorest/config"
)

func ConnectDB(config *config.Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error when connecting to DB:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Cannot reach the database:", err)
	}

	return db
}
