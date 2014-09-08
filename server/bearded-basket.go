package main

import (
	"fmt"
	"net/http"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/handlers"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	database.Init("prod")
	defer database.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/pdv", handlers.PDVsHandler)
	r.HandleFunc("/pdv/{id:[0-9]+}", handlers.PDVHandler)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
