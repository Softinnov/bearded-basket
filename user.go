package main

import (
	"net/http"
	"strconv"

	"github.com/Softinnov/bearded-basket/utils"
	"github.com/gorilla/mux"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

var users []User = []User{
	{1, "toto", "male"},
	{2, "titi", "female"},
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var user User
	for _, user = range users {
		if user.Id == id {
			break
		}
	}
	utils.WriteJSON(w, user)
}
