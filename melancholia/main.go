package main

import (
  "github.com/gerep/melancholia"
  "log"
  "net/http"
)

func main() {
  melancholia.CreateRoutes()

  log.Println("Listening on :4242")
  http.ListenAndServe(":4242", nil)
}
