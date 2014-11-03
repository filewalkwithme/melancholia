package main

import (
  "fmt"
  "net/http"
  "log"
)

type ApiFunc func(w http.ResponseWriter, r *http.Request)

func main() {
  createRoutes()

  log.Println("Listening on :4242")
  http.ListenAndServe(":4242", nil)
}

func createRoutes() {
  m := map[string]map[string]ApiFunc{
    "GET": {
      "/ping": ping,
      "/pong": pong,
    },
  }

  for method, routes := range m {
    for route, function := range routes {
      log.Printf("Creating route %s with method %s", route, method)
      http.HandleFunc(route, function)
    }
  }
}

func ping(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "I'm ping @ %s", r.URL.Path[1:])
}

func pong(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "I'm pong @ %s", r.URL.Path[1:])
}
