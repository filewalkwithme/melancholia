package router

import (
	"net/http"
	"encoding/json"

	"github.com/gerep/melancholia/models"
)

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