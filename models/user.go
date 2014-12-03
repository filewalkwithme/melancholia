package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	DB *sql.DB
	ID, Name, Email, Password string
}

func (u User) Save() (user User, err error) {
	var id int
	err = u.DB.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id", u.Name, u.Email, u.Password).Scan(&id)
	if err != nil {
		panic(err)
	}
	err = u.DB.QueryRow("SELECT id, email, name FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		panic(err)
	}
	return user, nil
}