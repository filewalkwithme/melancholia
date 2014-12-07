package models

import (
	"database/sql"
	"errors"

	"github.com/gerep/melancholia/libs"
	_ "github.com/lib/pq"
)

type User struct {
	DB *sql.DB
	ID, Name, Email, Password string
}

func (u User) Save() (user User, err error) {

	if libs.MinSize(u.Name, 3) != true {
		return user, errors.New(`{"error":"Name is too short"}`)
	}
	if libs.MaxSize(u.Name, 50) != true {
		return user, errors.New(`{"error":"Name is too long"}`)
	}

	if libs.MinSize(u.Email, 5) != true {
		return user, errors.New(`{"error":"E-mail is too short"}`)
	}
	if libs.MaxSize(u.Email, 50) != true {
		return user, errors.New(`{"error":"Email is too long"}`)
	}
	if u.Unique() != true {
		return user, errors.New(`{"error":"Email already used"}`)
	}

	var id int
	err = u.DB.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id", u.Name, u.Email, u.Password).Scan(&id)
	if err != nil {
		return user, err
	}

	err = u.DB.QueryRow("SELECT id, email, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Email, &user.Name)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u User) Unique() (result bool) {
	var id int
	err := u.DB.QueryRow("SELECT id FROM users WHERE email = $1 LIMIT 1", u.Email).Scan(&id)
	if err == sql.ErrNoRows {
		return true
	} else {
		return false
	}
	return true
}
