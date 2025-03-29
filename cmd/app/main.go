package main

import (
	"log"
	"net/http"
	"github.com/samvel333/gorest/config"
)

func main() {
	config := config.LoadConfig()

	mux := http.NewServeMux()
	log.Println("Server started at port", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, mux))
}
