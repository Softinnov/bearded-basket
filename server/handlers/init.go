package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

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

	rand.Seed(time.Now().UnixNano())

	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, 7)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	r.HandleFunc("/api/version", func(w http.ResponseWriter, h *http.Request) {
		fmt.Fprintf(w, "Version 1.0.0 %s", string(b))
	}).Methods("GET")

	http.Handle("/", r)
}
