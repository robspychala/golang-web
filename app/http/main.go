package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"

	"app/shared"
)

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("postgres", shared.GetDBUrl())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("started " + os.Args[0])
	log.Println("using " + shared.GetDBUrl())
	log.Println("environment " + os.Getenv("APP_ENV"))

	if shared.IsDevEnvironment() {
		log.Println("url http://" + shared.GetIPAddress() + ":" + os.Getenv("APP_PORT") + "/")
	}

	log.Fatal(http.ListenAndServe(":"+os.Getenv("APP_PORT"), NewHttpHandler()))

}
