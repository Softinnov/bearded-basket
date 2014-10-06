package handlers

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/mux"
)

func Init(c *utils.Context) {

	r := mux.NewRouter()

	r.Handle("/api/users/{id:[0-9]+}", authHandler{c, editUser}).Methods("PUT")
	r.Handle("/api/users", authHandler{c, indexUsers}).Methods("GET")

	http.Handle("/", r)
}
