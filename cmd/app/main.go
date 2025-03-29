package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	var port int = 5000
	log.Println("Сервер запущен на порту", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
