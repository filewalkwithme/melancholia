package router

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gerep/melancholia/models"
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
	rows, err := r.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	users := make([]models.User, 0)

	var (
		id int
		name, email string
	)

	for rows.Next() {
		rows_err := rows.Scan(&id, &name, &email)
		if rows_err != nil {
			panic(rows_err)
		}

		users = append(users, models.User{ID: strconv.Itoa(id), Name: name, Email: email})
	}

	if len(users) > 0 {
		json.NewEncoder(w).Encode(users)
	} else {
		json.NewEncoder(w).Encode(`{}`)
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