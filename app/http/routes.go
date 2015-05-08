package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"app/shared"
)

func NewHttpHandler() http.Handler {

	handler := mux.NewRouter().StrictSlash(true)
	handler.HandleFunc("/", IndexPageHandler).Methods("GET")
	handler.HandleFunc("/api/hi", ApiHelloPageHandler).Methods("GET")

	if shared.IsDevEnvironment() {
		handler.Handle("/asset/{rest}", http.StripPrefix("/asset/", http.FileServer(http.Dir("../http-client-asset"))))
	}

	return handler
}
