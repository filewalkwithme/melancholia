package main

import (
	"github.com/gerep/melancholia"
	"net/http"
)

func main() {
	melancholia.CreateRoutes()
	http.ListenAndServe(":4242", nil)
}
