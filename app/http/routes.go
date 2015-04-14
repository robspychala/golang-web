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

  if os.Getenv("APP_ENV") == "dev" {
    handler.Handle("/asset/{rest}", http.StripPrefix("/asset/", http.FileServer(http.Dir("../http-client-asset"))))
  }

  return handler
}
