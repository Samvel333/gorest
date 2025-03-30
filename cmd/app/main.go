package main

import (
	"log"
	"net/http"

	_ "github.com/samvel333/gorest/cmd/app/docs" // Adjust this based on your module path
	"github.com/samvel333/gorest/config"
	"github.com/samvel333/gorest/internal/handlers"
	"github.com/samvel333/gorest/internal/repository"
	"github.com/samvel333/gorest/internal/services"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title People API
// @version 1.0
// @description People API (test)
// @BasePath /
func main() {
	config := config.LoadConfig()
	// DB Connecting
	db := services.ConnectDB(config)

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
