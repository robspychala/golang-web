package main

import (
    "log"
    "encoding/json"
    "net/http"

    // "app/shared"
)

func ApiHelloPageHandler(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)
  if err := json.NewEncoder(w).Encode(map[string]interface{}{ "success": true}); err != nil {
    log.Fatal(err)
  }
}
