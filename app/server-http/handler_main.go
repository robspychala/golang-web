package main

import (
    "os"
    "log"
    "net/http"
    "html/template"

    "app/shared"
)

var indexTempl = template.Must(template.ParseFiles("template/index.html"))

func IndexPageHandler(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  err := indexTempl.Execute(w, map[string]string{"Env": os.Getenv("APP_ENV")})

  if err != nil {
    log.Fatal(err)
  }
}
