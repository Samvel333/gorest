package main

import (
	"log"
	"net/http"

	_ "github.com/samvel333/gorest/cmd/app/docs"
	"github.com/samvel333/gorest/config"
	"github.com/samvel333/gorest/internal/handlers"
	"github.com/samvel333/gorest/internal/repository"
	"github.com/samvel333/gorest/internal/services"
	httpSwagger "github.com/swaggo/http-swagger"
)


// @title People API (Task)
// @version 1.0
// @description A simple API for managing people data.
// @BasePath /
// @contact.name Samvel Sadoyan
// @contact.email sadoyansamvel@yandex.com
// @host localhost:7000
// @schemes http
func main() {
	config := config.LoadConfig()
	// DB Connecting
	db := services.ConnectDB(config)

	repo := repository.NewRepository(db)
	handler := handlers.NewHandler(repo)

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", httpSwagger.WrapHandler)

	mux.HandleFunc("POST /person", handler.CreatePersonHandler)
	mux.HandleFunc("GET /people", handler.GetPeopleHandler)
	mux.HandleFunc("GET /person", handler.GetPersonByIDHandler)
	mux.HandleFunc("DELETE /person", handler.DeletePersonHandler)
	mux.HandleFunc("PUT /person", handler.UpdatePersonHandler)

	log.Println("Server started at port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, mux))
}
