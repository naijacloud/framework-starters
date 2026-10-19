package main

import (
	"net/http"
	"os"

	"example.com/lib"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte(lib.Msg())) })
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
