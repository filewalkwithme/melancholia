package models

import (
		"testing"
		"database/sql"
		_ "github.com/lib/pq"
)

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=melancholia_test sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	return db
}

func eraseDB() {
	db := setupDB()
	db.Exec("TRUNCATE users")
}

func user() User{
	return User{Name: "Jaspion", Email: "jaspion@daileon", Password: "123456", DB: setupDB()}
}

func init() {
	eraseDB()
}

func TestNameTooShort(t *testing.T) {
	user := user()
	user.Name = "Jas"
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Name is too short"}` {
			t.Errorf("Name validation failed: ", err.Error())
	}
}

func TestNameTooLong(t *testing.T) {
	user := user()
	user.Name = "Jaspion and Daileon together will destroy Satan Goss forever!"
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Name is too long"}` {
		t.Errorf("Name validation failed: ", err.Error())
	}
}

func TestEmailTooShort(t *testing.T) {
	user := user()
	user.Email = "a@a."
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Email is too short"}` {
		t.Errorf("Email validation failed: ", err.Error())
	}
}

func TestEmailTooLong(t *testing.T) {
	user := user()
	user.Email = "the_powerfull_jaspion_will_destroy@satan_goss.com.br"
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Email is too long"}` {
		t.Errorf("Email validation failed: ", err.Error())
	}
}

func TestEmailNotValid(t *testing.T) {
	user := user()
	user.Email = "this is @not an email.com.br"
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Email is not valid"}` {
		t.Errorf("Email validation failed: ", err.Error())
	}
}

func TestEmailIsUnique(t *testing.T) {
	jaspion := user()
	jaspion.Save()

	jiban := user()
	_, err := jiban.Save()

	if err != nil && err.Error() != `{"error":"Email is taken"}` {
		t.Errorf("Email validation failed: ", err.Error())
	}
}

func TestAuthentication(t *testing.T) {
	jaspion := user()
	jaspion.Save()

	_, err := jaspion.Authenticate()
	if err != nil {
		t.Errorf("User authentication failed: ", err.Error())
	}
}