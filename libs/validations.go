package libs

import (
	"database/sql"
	"net/mail"

	_ "github.com/lib/pq"
)

type Validation struct {
	Error string
	OK    bool
}

func (v Validation) Message(message string) (string, bool) {
	if v.OK == false {
		return message, v.OK
	}
	return "", v.OK
}

func (v Validation) MinSize(who string, size int) Validation {
	if len(who) < size {
		return Validation{OK: false}
	}
	return Validation{OK: true}
}

func (v Validation) MaxSize(who string, size int) Validation {
	if len(who) > size {
		return Validation{OK: false}
	}
	return Validation{OK: true}
}

func (v Validation) Email(email string) Validation {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return Validation{OK: false}
	}
	return Validation{OK: true}
}

func (v Validation) Unique(what string, from string, value string, db *sql.DB) Validation {
	var id int
	query := "SELECT id FROM " + from + " WHERE " + what + " = '" + value + "'"
	err := db.QueryRow(query).Scan(&id)
	if err == sql.ErrNoRows {
		return Validation{OK: true}
	} else {
		return Validation{OK: false}
	}
	return Validation{OK: true}
}
