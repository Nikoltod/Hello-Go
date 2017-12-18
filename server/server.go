package main

import (
	"io"
	"net/http"
)

func ws(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", ws)
	http.ListenAndServe(":8080", nil)
}
