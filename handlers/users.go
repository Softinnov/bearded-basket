package handlers

import (
	"net/http"
	"strconv"

	"github.com/Softinnov/bearded-basket/models"
	"github.com/Softinnov/bearded-basket/utils"
	"github.com/gorilla/mux"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := models.GetUsers()
	utils.WriteJSON(w, users)
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, _ := models.GetUser(id)
	utils.WriteJSON(w, user)
}
