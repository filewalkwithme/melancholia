package router

import (
	"encoding/json"
	"net/http"

	"github.com/maiconio/melancholia/models"
)

func (r Router) createUser(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	email := req.FormValue("email")
	password := req.FormValue("password")

	user := models.User{Name: name, Email: email, Password: password, DB: r.DB}
	result, err := user.Save()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(result)
	}
}

func (r Router) authenticate(w http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")

	user := models.User{Email: email, Password: password, DB: r.DB}
	result, err := user.Authenticate()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(result)
	}
}
