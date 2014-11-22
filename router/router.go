package router

import (
	"net/http"
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
)

type ApiFunc func(w http.ResponseWriter, req *http.Request)

type Router struct {
	DB *sql.DB
}

func (r Router) CreateRoutes() {

	gmux := mux.NewRouter()

	http.Handle("/", gmux)

	m := map[string]map[string]ApiFunc{
		"GET": {
			"/users": r.users,
		},
		"POST": {
			"/users": r.createUser,
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

func (r Router) users(w http.ResponseWriter, req *http.Request) {
	var (
		id int
		name string
	)
	rows, err := r.DB.Query("SELECT id, name FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows_err := rows.Scan(&id, &name)
		if rows_err != nil {
			panic(rows_err)
		}
		fmt.Println(id, name)
	}
}

func (r Router) createUser(w http.ResponseWriter, req *http.Request) {
}

func updateUser(w http.ResponseWriter, req *http.Request) {
	//
}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	//
}