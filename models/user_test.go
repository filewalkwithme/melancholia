package models

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
)

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "user=docker password='docker' dbname=docker sslmode=disable")
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
	createScript :=
		`DROP TABLE IF EXISTS "public"."users";
CREATE TABLE "public"."users" (
	id serial primary key, -- Sequence structure is created automatically when using 'serial'
	name varchar(40) NOT NULL COLLATE "default",
	email varchar(40) NOT NULL COLLATE "default",
	password varchar(60)
)
WITH (OIDS=FALSE);`
	db.Exec(createScript)

	db.Exec("TRUNCATE users")

}

func user() User {
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
		t.Errorf("Name validation failed: %v", err.Error())
	}
}

func TestNameTooLong(t *testing.T) {
	user := user()
	user.Name = "Jaspion and Daileon together will destroy Satan Goss forever!"
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Name is too long"}` {
		t.Errorf("Name validation failed: %v", err.Error())
	}
}

func TestPasswordTooShort(t *testing.T) {
	user := user()
	user.Password = "123"
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Password is too short"}` {
		t.Errorf("Password validation failed: %v", err.Error())
	}
}

func TestEmailTooShort(t *testing.T) {
	user := user()
	user.Email = "a@a."
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Email is too short"}` {
		t.Errorf("Email validation failed: %v", err.Error())
	}
}

func TestEmailTooLong(t *testing.T) {
	user := user()
	user.Email = "the_powerfull_jaspion_will_destroy@satan_goss.com.br"
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Email is too long"}` {
		t.Errorf("Email validation failed: %v", err.Error())
	}
}

func TestEmailNotValid(t *testing.T) {
	user := user()
	user.Email = "this is @not an email.com.br"
	_, err := user.Save()

	if err != nil && err.Error() != `{"error":"Email is not valid"}` {
		t.Errorf("Email validation failed: %v", err.Error())
	}
}

func TestEmailIsUnique(t *testing.T) {
	jaspion := user()
	jaspion.Save()

	jiban := user()
	_, err := jiban.Save()

	if err != nil && err.Error() != `{"error":"Email is taken"}` {
		t.Errorf("Email validation failed: %v", err.Error())
	}
}

func TestAuthentication(t *testing.T) {
	jaspion := user()
	jaspion.Save()

	_, err := jaspion.Authenticate()
	if err != nil {
		t.Errorf("User authentication failed: %v", err.Error())
	}
}

func TestAuthenticationWithUnsavedUser(t *testing.T) {
	invalidUser := user()
	invalidUser.Email = "invalid@user.com.br"

	_, err := invalidUser.Authenticate()
	if err == nil {
		t.Errorf("User authentication failed: Unsaved user authorized")
	}
}

func TestAuthenticationWithInvalidPassword(t *testing.T) {
	jaspion := user()
	jaspion.Save()

	jaspion.Password = "wrong password"

	_, err := jaspion.Authenticate()
	if err == nil {
		t.Errorf("User authentication failed: User with wrong password authorized")
	}
}
