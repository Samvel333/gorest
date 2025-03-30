package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/samvel333/gorest/cmd/app/docs" // Adjust this based on your module path
	"github.com/samvel333/gorest/config"
	"github.com/samvel333/gorest/internal/handlers"
	"github.com/samvel333/gorest/internal/repository"
	httpSwagger "github.com/swaggo/http-swagger"
)

func mustConnectDB(config *config.Config) *sql.DB {
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

// @title People API
// @version 1.0
// @description People API (test)
// @BasePath /
func main() {
	config := config.LoadConfig()
	// DB Connecting
	db := mustConnectDB(config)

	repo := repository.NewRepository(db)
	handler := handlers.NewHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	mux.HandleFunc("POST /people", handler.CreatePersonHandler)

	mux.HandleFunc("GET /people", handler.GetPeopleHandler)
	mux.HandleFunc("GET /person", handler.GetPersonByIDHandler)
	mux.HandleFunc("DELETE /people/delete", handler.DeletePersonHandler)
	mux.HandleFunc("PUT /people/update", handler.UpdatePersonHandler)

	log.Println("Server started at port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, mux))
}
