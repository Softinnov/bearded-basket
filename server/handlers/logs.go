package handlers

import (
	"log"
	"net/http"

	"github.com/Softinnov/bearded-basket/server/utils"
)

func CheyenneLogHTTP(msg string, status int) {
	log.Printf("Cheyenne %d: %s\n", status, msg)
}

func CheyenneErrorHTTP(w http.ResponseWriter, err string, status int) {
	log.Printf("Cheyenne %d: %s\n", status, err)
	http.Error(w, http.StatusText(status), http.StatusUnauthorized)
}

func LogHTTP(e *utils.SError, w http.ResponseWriter, r *http.Request) {
	if e != nil {
		log.Printf("HTTP %s %d: %s | %s\n", r.Method, e.Status, e.Back, r.URL)
		if e.Front != nil {
			http.Error(w, e.Front.Error(), e.Status)
		}
		http.Error(w, "", e.Status)
	} else {
		log.Printf("HTTP %s %d: %s %s\n", r.Method, http.StatusOK,
			http.StatusText(http.StatusOK), r.URL)
	}
}
