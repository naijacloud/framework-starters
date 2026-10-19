package handlers

import "net/http"

func Hello(w http.ResponseWriter, r *http.Request) { w.Write([]byte("hello from go stdlib")) }
