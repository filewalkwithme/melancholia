package melancholia

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type ApiFunc func(w http.ResponseWriter, r *http.Request)

func CreateRoutes() {

	gmux := mux.NewRouter()

	http.Handle("/", gmux)

	m := map[string]map[string]ApiFunc{
		"GET": {
			"/ping": ping,
			"/pong": pong,
		},
		"POST": {
			"/authenticate": authenticate,
		},
	}

	for method, routes := range m {
		for route, function := range routes {
			log.Printf("Creating route %s with method %s", route, method)
			gmux.HandleFunc(route, function).Methods(method)
		}
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm ping @ %s", r.URL.Path[1:])
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm pong @ %s", r.URL.Path[1:])
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm authenticate @ %s", r.URL.Path[1:])
}
