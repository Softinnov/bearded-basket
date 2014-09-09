package main

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.Init("prod")
	defer database.Close()

	r := mux.NewRouter()

	r.HandleFunc("/pdv", handlers.PDVsHandler)
	r.HandleFunc("/pdv/{id:[0-9]+}", handlers.PDVHandler)

	http.Handle("/", r)

	http.ListenAndServe(":8000", nil)
}
