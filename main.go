package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// declare a new router
	r := mux.NewRouter()

	r.HandleFunc("/hello", handler).Methods("GET")

	// instead of
	// http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// takes a “writer” as its first argument.
	fmt.Fprintf(w, "Hello From Mars!!!")
}
