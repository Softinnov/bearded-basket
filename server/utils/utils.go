package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	Unauthorised = iota + 1
	Internal
)

type MError struct {
	Id  int
	Err error
}

func (e *MError) Error() string {
	return e.Err.Error()
}

func (e *MError) HTTPError() (int, string) {
	switch e.Id {
	case Unauthorised:
		return http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)
	default:
		return http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)
	}
}

func NewMError(id int, f string, a ...interface{}) *MError {
	return &MError{id, fmt.Errorf(f, a...)}
}

func WriteJSON(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("content-type", "application/json; charset=utf-8")
	_, err = w.Write(data)
	return err
}

func BuildSqlSets(b []byte) (string, error) {
	var buf bytes.Buffer
	var data map[string]interface{}

	err := json.Unmarshal(b, &data)
	if err != nil {
		return "", err
	}
	flag := false
	for key, val := range data {
		if flag {
			buf.WriteString(", ")
		}
		switch val.(type) {
		case string:
			buf.WriteString(key + "=" + fmt.Sprintf("%q", val))
		default:
			buf.WriteString(key + "=" + fmt.Sprintf("%v", val))
		}
		flag = true
	}
	return buf.String(), nil
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
