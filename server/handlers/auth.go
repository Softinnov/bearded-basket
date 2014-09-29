package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Softinnov/bearded-basket/server/models"
	"github.com/Softinnov/bearded-basket/server/utils"
)

var (
	COOKIE_NAME = "RSPSID"
	URL         = "http://localhost:8000/pdv"
	URL_WS      = "/remote/ws.rsp"
)

type authHandler struct {
	c  *utils.Context
	fn func(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error)
}

func cookieAuth(c *utils.Context, w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	req, err := http.NewRequest("GET", URL+URL_WS, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	req.AddCookie(cookie)
	res, err := http.DefaultTransport.RoundTrip(req) // Avoid redirection
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}
	switch res.StatusCode {
	case 200:
		log.Println("200 OK")
		session := &models.Session{}
		err := json.NewDecoder(res.Body).Decode(session)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return false
		}
		res.Body.Close()
		if err := models.StoreInCookies(c.Store, session, w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return false
		}
	case 302:
		url, _ := res.Location()
		log.Println("302, ", url.String())
		http.Error(w, url.String(), http.StatusNotFound)
		return false
	default:
		log.Println("bad request")
		http.Error(w, "bad request", http.StatusInternalServerError)
		return false
	}
	return true
}

func (ah authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !cookieAuth(ah.c, w, r) {
		return
	}
	status, err := ah.fn(ah.c, w, r)
	if err != nil {
		log.Printf("HTTP %d: %q\n", status, err)
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}
