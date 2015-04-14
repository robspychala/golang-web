package main

import (
    "os"
    "net/http"

    "github.com/gorilla/mux"
)

func SetupHttpHandler() http.Handler {

  handler := mux.NewRouter().StrictSlash(true)
  handler.HandleFunc("/", IndexPageHandler).Methods("GET")
  handler.HandleFunc("/api/hi", ApiHelloPageHandler).Methods("GET")

  if os.Getenv("APP_PORT") == "dev" {
    handler.Handle("/asset", http.FileServer(http.Dir("../client-asset")))
  }

  return handler
}
