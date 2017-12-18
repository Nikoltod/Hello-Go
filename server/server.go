package main

import (
	"io"
	"net/http"
)

//Simple handler for incoming requests using the 'io' package
func ws(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

//Handling the requests on '/' using port 8080
func main() {
	http.HandleFunc("/", ws)
	http.ListenAndServe(":8080", nil)
}
