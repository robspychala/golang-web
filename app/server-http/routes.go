package main

import (
    "log"
    "os"
    "net/http"
    "github.com/gorilla/mux"
    "database/sql"
    _ "github.com/lib/pq"
)

func SetupHttpHandler() http.Handler {

  handler := mux.NewRouter().StrictSlash(true)
  handler.HandleFunc("/", IndexPageHandler).Methods("GET")
  handler.HandleFunc("/api/hi", ApiHelloPageHandler).Methods("GET")

  return handler
}
