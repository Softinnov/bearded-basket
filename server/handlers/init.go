package handlers

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/mux"
)

func Init(c *utils.Context) {

	r := mux.NewRouter()

	// AUTHENTICATED
	r.Handle("/api/users/{id:[0-9]+}", authHandler{c, editUser}).Methods("PUT")
	r.Handle("/api/users", authHandler{c, indexUsers}).Methods("GET")
	r.Handle("/api/users", authHandler{c, newUser}).Methods("POST")
	r.Handle("/api/users/{id:[0-9]+}", authHandler{c, deleteUser}).Methods("DELETE")

	// NOT AUTHENTICATED
	r.Handle("/api/roles", basicHandler{c, indexRoles}).Methods("GET")

	http.Handle("/", r)
}
