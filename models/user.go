package models

import (
	"database/sql"
	"errors"

	"github.com/gerep/melancholia/libs"
	_ "github.com/lib/pq"
)

type User struct {
	DB                        *sql.DB
	ID, Name, Email, Password string
}

func (u User) Save() (User, error) {

	v := libs.Validation{}

	if msg, err := v.MinSize(u.Name, 5).Message(`{"error":"Name is too short"}`); err != true {
		return u, errors.New(msg)
	}

	if msg, err := v.MaxSize(u.Name, 40).Message(`{"error":"Name is too long"}`); err != true {
		return u, errors.New(msg)
	}

	if msg, err := v.MinSize(u.Email, 5).Message(`{"error":"Email is too short"}`); err != true {
		return u, errors.New(msg)
	}

	if msg, err := v.MaxSize(u.Email, 40).Message(`{"error":"Email is too long"}`); err != true {
		return u, errors.New(msg)
	}

	if msg, err := v.Email(u.Email).Message(`{"error":"Email is not valid"}`); err != true {
		return u, errors.New(msg)
	}

	if msg, err := v.Unique("email", "users", u.Email, u.DB).Message(`{"error":"Email is taken"}`); err != true {
		return u, errors.New(msg)
	}

	var id int
	err := u.DB.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id", u.Name, u.Email, u.Password).Scan(&id)
	if err != nil {
		return u, err
	}

	err = u.DB.QueryRow("SELECT id, email, name FROM users WHERE id = $1", id).Scan(&u.ID, &u.Email, &u.Name)
	if err != nil {
		return u, err
	}
	return u, nil
}
