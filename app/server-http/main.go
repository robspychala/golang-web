package main

import (
    "log"
    "os"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
)

func main() {

    var err error
    db, err = sql.Open("postgres", os.Getenv("APP_DBURL"))
    if err != nil {
      log.Fatal(err)
    }

    log.Print("started " + os.Args[0])

    handler := SetupRouter()

    log.Fatal(http.ListenAndServe(":" + os.Getenv("APP_PORT"), handler))

}
