package main

import (
	"net/http"
	"os"

	"example.com/api/internal/greet"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(greet.Msg())) })
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
