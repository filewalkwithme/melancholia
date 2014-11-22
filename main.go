package main

import (
	"net/http"
	"database/sql"

	"github.com/gerep/melancholia/router"
	_ "github.com/lib/pq"
)

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=melancholia sslmode=disable")
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

func main() {
	r := router.Router{DB: setupDB()}
	r.CreateRoutes()
	http.ListenAndServe(":4242", nil)
}