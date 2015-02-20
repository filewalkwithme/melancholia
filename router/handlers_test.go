package router

import (
	"database/sql"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
)

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "user=melancholia password='m1e2l3a4' dbname=melancholia_test sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
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
