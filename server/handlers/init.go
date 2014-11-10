package handlers

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/mux"
)

func Init(c *utils.Context) {

	r := mux.NewRouter()

	// NEED AUTHENTICATION
	r.Handle("/api/users/{id:[0-9]+}", authHandler{c, editUser}).Methods("PUT")
	r.Handle("/api/users", authHandler{c, indexUsers}).Methods("GET")
	r.Handle("/api/user", authHandler{c, getCurrentUser}).Methods("GET")
	r.Handle("/api/users", authHandler{c, newUser}).Methods("POST")
	r.Handle("/api/users/{id:[0-9]+}", authHandler{c, deleteUser}).Methods("DELETE")

	// NO AUTHENTICATION
	r.Handle("/api/roles", basicHandler{c, indexRoles}).Methods("GET")

	http.Handle("/", r)
}
