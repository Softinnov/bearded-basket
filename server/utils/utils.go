package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	_, err = w.Write(data)
	return err
}

func CheyenneLogHTTP(msg string, status int) {
	log.Printf("Cheyenne %d: %s\n", status, msg)
}

func CheyenneErrorHTTP(w http.ResponseWriter, err string, status int) {
	log.Printf("Cheyenne %d: %s\n", status, err)
	http.Error(w, http.StatusText(status), http.StatusUnauthorized)
}

func LogHTTP(w http.ResponseWriter, err error, status int, r *http.Request) {
	if err != nil {
		log.Printf("HTTP %s %d: %s %s\n", r.Method, status, err, r.URL)
		http.Error(w, http.StatusText(status), http.StatusUnauthorized)
	} else {
		log.Printf("HTTP %s %d: %s %s\n", r.Method, status, http.StatusText(status), r.URL)
	}
}
