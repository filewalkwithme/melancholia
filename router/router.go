package router

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gerep/melancholia/models"
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
			"/users/{id:[0-9]+}": r.getUser,
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

func (r Router) getUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	user := models.User{DB: r.DB, Id: id}
	result, err := user.Get()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func (r Router) createUser(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	email := req.FormValue("email")
	password := req.FormValue("password")

	user := models.User{Name: name, Email: email, Password: password, DB: r.DB}
	result, err := user.Save()

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(result)
	}
}

func updateUser(w http.ResponseWriter, req *http.Request) {
	//
}

func deleteUser(w http.ResponseWriter, req *http.Request) {
	//
}
