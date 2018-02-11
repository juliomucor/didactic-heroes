package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// takes a “writer” as its first argument.
	fmt.Fprintf(w, "Hello From Mars!!!")
}

// refactor router to test
// as a pattern: Once we’ve separated our route constructor function, let’s test our routing:
func newRouter() *mux.Router {
	// declare a new router
	r := mux.NewRouter()
	// instead of
	// http.HandleFunc("/", handler)
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}
