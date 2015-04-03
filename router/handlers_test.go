package router

import (
	"database/sql"
	"net/http"
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

	return db
}

func init() {
	testRouter := Router{DB: setupDB()}
	muxer := testRouter.CreateRoutes()

	go http.ListenAndServe(":4242", muxer)
}

func TestPostAuthenticate(t *testing.T) {
	resp, err := http.PostForm("http://localhost:4242/authenticate", nil)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 500 {
		t.Errorf("POST /authenticate failed: Expected 500, received %d", resp.StatusCode)
	}
}

func TestPostUsers(t *testing.T) {
	resp, err := http.PostForm("http://localhost:4242/users", nil)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 500 {
		t.Errorf("POST /users failed: Expected 500, received %d", resp.StatusCode)
	}
}
