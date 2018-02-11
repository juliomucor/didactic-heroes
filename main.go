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

	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")

	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix(assets, http.FileServer(staticFileDirectory))

	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix(assets).Handler(staticFileHandler).Methods("GET")
	return r
}

const assets = "/assets/"
