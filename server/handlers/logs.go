package handlers

import (
	"log"
	"net/http"

	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/fatih/color"
)

func CheyenneLogHTTP(msg string, status int) {
	color.Set(color.FgYellow)
	log.Printf("Cheyenne %d: %s\n", status, msg)
	color.Unset()
}

func CheyenneErrorHTTP(w http.ResponseWriter, err string, status int) {
	color.Set(color.FgRed)
	log.Printf("Cheyenne %d: %s\n", status, err)
	color.Unset()
	http.Error(w, http.StatusText(status), http.StatusUnauthorized)
}

func LogHTTP(e *utils.SError, w http.ResponseWriter, r *http.Request) {
	if e != nil {
		color.Set(color.FgRed)
		log.Printf("HTTP %s %d: %s | %s\n", r.Method, e.Status, e.Back, r.URL)
		color.Unset()
		if e.Front != nil {
			http.Error(w, e.Front.Error(), e.Status)
		}
		http.Error(w, "", e.Status)
	} else {
		log.Printf("HTTP %s %d: %s %s\n", r.Method, http.StatusOK,
			http.StatusText(http.StatusOK), r.URL)
	}
}
