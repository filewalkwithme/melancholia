package main

import (
	"database/sql"
	"net/http"

	"github.com/gerep/melancholia/router"
	_ "github.com/lib/pq"
)

func setupDB() *sql.DB {
	db, err := sql.Open("postgres", "user=melancholia password='m1e2l3a4' dbname=melancholia sslmode=disable")
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
