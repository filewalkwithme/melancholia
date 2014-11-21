package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ApiFunc func(w http.ResponseWriter, r *http.Request)

func CreateRoutes() {

	gmux := mux.NewRouter()

	http.Handle("/", gmux)

	m := map[string]map[string]ApiFunc{
		"GET": {
			"/users": users,
		},
		"POST": {
			"/users": createUser,
		},
		"DELETE": {
			"/users": deleteUser,
		},
		"PUT": {
			"/users": updateUser,
		},
	}

	for method, routes := range m {
		for route, function := range routes {
			gmux.HandleFunc(route, function).Methods(method)
		}
	}
}

func users(w http.ResponseWriter, r *http.Request) {
	// 
}

func createUser(w http.ResponseWriter, r *http.Request) {
	//
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	//
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	//
}