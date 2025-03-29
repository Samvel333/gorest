package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file")
	}

	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	log.Println("Server started at port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
