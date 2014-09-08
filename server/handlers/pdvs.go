package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Softinnov/bearded-basket/server/models"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/mux"
)

func PDVsHandler(w http.ResponseWriter, r *http.Request) {
	pdvs, err := models.GetPDVs()
	if err != nil {
		log.Println("PDVsHandler:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, pdvs)
}

func PDVHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("PDVHandler:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pdv, _ := models.GetPDV(id)
	utils.WriteJSON(w, pdv)
}
