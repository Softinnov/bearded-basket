package handlers

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, s int, v interface{}) error {
	d, e := json.Marshal(v)
	if e != nil {
		return e
	}
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(s)
	_, e = w.Write(d)
	if e != nil {
		return e
	}
	return nil
}
