package main

import (
	"net/http"
	"os"

	"example.com/stdlib/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handlers.Hello)
	http.ListenAndServe(":"+port, nil)
}
