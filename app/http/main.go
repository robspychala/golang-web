package main

import (
    "log"
    "os"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"

    "app/shared"
)

var db *sql.DB

func main() {

    var err error
    db, err = sql.Open("postgres", os.Getenv("APP_DBURL"))
    if err != nil {
      log.Fatal(err)
    }

    log.Println("started " + os.Args[0])
    log.Println("using " + os.Getenv("APP_DBURL"))
    log.Println("environment " + os.Getenv("APP_ENV"))
    log.Println("url http://" + shared.GetIPAddress() + ":" + os.Getenv("APP_PORT") + "/")

    handler := SetupHttpHandler()

    log.Fatal(http.ListenAndServe(":" + os.Getenv("APP_PORT"), handler))

}
