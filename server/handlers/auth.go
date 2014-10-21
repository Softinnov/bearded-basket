package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Softinnov/bearded-basket/server/utils"
)

var (
	COOKIE_NAME = "RSPSID"
	URL         = "/pdv"
	URL_WS      = "/remote/ws.rsp"
)

type authHandler struct {
	c  *utils.Context
	fn func(c *utils.Context, w http.ResponseWriter, r *http.Request) (int, error)
}

func cookieAuth(c *utils.Context, w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		utils.LogHTTP(w, err, http.StatusUnauthorized, r)
		return false
	}
	req, err := http.NewRequest("GET", *c.Chey+URL+URL_WS, nil)
	if err != nil {
		utils.LogHTTP(w, err, http.StatusInternalServerError, r)
		return false
	}
	req.AddCookie(cookie)
	log.Printf("%#v\n", cookie)
	res, err := http.DefaultTransport.RoundTrip(req) // Avoid redirection
	if err != nil {
		utils.LogHTTP(w, err, http.StatusInternalServerError, r)
		return false
	}
	switch res.StatusCode {
	case 200:
		utils.CheyenneLogHTTP(r.URL.String(), res.StatusCode)
		session := &utils.Session{}
		err := json.NewDecoder(res.Body).Decode(session)
		if err != nil {
			utils.LogHTTP(w, err, http.StatusInternalServerError, r)
			return false
		}
		res.Body.Close()
		if err := utils.StoreInCookies(c.Store, session, w, r); err != nil {
			utils.LogHTTP(w, err, http.StatusInternalServerError, r)
			return false
		}
	case 302:
		url, _ := res.Location()
		utils.CheyenneErrorHTTP(w, url.String(), res.StatusCode)
		return false
	default:
		utils.CheyenneErrorHTTP(w, "", res.StatusCode)
		return false
	}
	return true
}

func (ah authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !cookieAuth(ah.c, w, r) {
		return
	}

	s, err := utils.SessionFromCookies(ah.c.Store, r)
	if err != nil || s == nil {
		utils.LogHTTP(w, err, http.StatusUnauthorized, r)
		return
	}
	if s.Role < 3 {
		utils.LogHTTP(w, errors.New("role unauthorized"), http.StatusUnauthorized, r)
		return
	}
	ah.c.Session = s
	status, err := ah.fn(ah.c, w, r)
	utils.LogHTTP(w, err, status, r)
}
