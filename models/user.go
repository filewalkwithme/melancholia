package models

import (
	"database/sql"
	"errors"

	"github.com/gerep/melancholia/libs"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	DB       *sql.DB `json:"-"`
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
}

func (u User) Save() (User, error) {

	v := libs.Validation{}

	if msg, err := v.MinSize(u.Name, 4).Message(`{"error":"Name is too short"}`); err != true {
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

	if msg, err := v.MinSize(u.Password, 5).Message(`{"error":"Password is too short"}`); err != true {
		return u, errors.New(msg)
	}

	if msg, err := v.Email(u.Email).Message(`{"error":"Email is not valid"}`); err != true {
		return u, errors.New(msg)
	}

	if msg, err := v.Unique("email", "users", u.Email, u.DB).Message(`{"error":"Email is taken"}`); err != true {
		return u, errors.New(msg)
	}

	password := []byte(u.Password)

	hashedpassword, bcrypt_err := bcrypt.GenerateFromPassword(password, 10)
	if bcrypt_err != nil {
		panic(bcrypt_err)
	}

	err := u.DB.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id", u.Name, u.Email, hashedpassword).Scan(&u.ID)
	if err != nil {
		return u, err
	}

	err = u.DB.QueryRow("SELECT id, email, name FROM users WHERE id = $1", u.ID).Scan(&u.ID, &u.Email, &u.Name)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (u User) Authenticate() (User, error) {
	var password []byte
	err := u.DB.QueryRow("SELECT id, name, password FROM users WHERE email = $1", u.Email).Scan(&u.ID, &u.Name, &password)
	if err == sql.ErrNoRows {
		return u, errors.New(`{"User not found"}`)
	}
	current_password := []byte(u.Password)
	err = bcrypt.CompareHashAndPassword(password, current_password)
	if err != nil {
		return u, errors.New(`{"User not found"}`)
	} else {
		return u, err
	}
	return u, nil
}
