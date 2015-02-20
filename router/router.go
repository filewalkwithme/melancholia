package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

type Handlers func(w http.ResponseWriter, req *http.Request)

type Router struct {
	DB *sql.DB
}

func (r Router) CreateRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	m := map[string]map[string]Handlers{
		"POST": {
			"/users":        r.createUser,
			"/authenticate": r.authenticate,
		},
	}

	for method, routes := range m {
		for route, function := range routes {
			router.HandleFunc(route, function).Methods(method)
		}
	}

	return router
}
