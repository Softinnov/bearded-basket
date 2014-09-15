package main

import (
	"flag"
	"net/http"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/handlers"
	"github.com/gorilla/mux"
)

var (
	db = flag.String("db", "root:@/prod",
		"database, usage: user:password@addr/dbname")
)

func main() {
	flag.Parse()

	database.Init(*db)
	defer database.Close()

	r := mux.NewRouter()

	r.HandleFunc("/pdv", handlers.PDVsHandler)
	r.HandleFunc("/pdv/{id:[0-9]+}", handlers.PDVHandler)

	http.Handle("/", r)

	http.ListenAndServe(":8000", nil)
}
