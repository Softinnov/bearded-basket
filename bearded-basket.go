package main

import (
	"fmt"
	"net/http"

	"github.com/Softinnov/bearded-basket/database"
	"github.com/Softinnov/bearded-basket/handlers"
	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	database.Init("bearded")
	defer database.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler)

	r.HandleFunc("/users", handlers.UsersHandler)
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UserHandler)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
