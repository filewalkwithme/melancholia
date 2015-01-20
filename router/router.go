package router

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gerep/melancholia/models"
	"github.com/gorilla/mux"
)

type ApiFunc func(w http.ResponseWriter, req *http.Request)

type Router struct {
	DB *sql.DB
}

func (r Router) CreateRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	m := map[string]map[string]ApiFunc{
		"POST": {
			"/users": r.createUser,
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

func (r Router) authenticate(w http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")

	user := models.User{Email: email, Password: password, DB: r.DB}
	result, err := user.Authenticate()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(result)
	}
}