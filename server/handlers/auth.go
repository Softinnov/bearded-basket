package handlers

import (
	"encoding/json"
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
	fn func(c *utils.Context, w http.ResponseWriter, r *http.Request) *utils.SError
}

func cookieAuth(c *utils.Context, w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		LogHTTP(&utils.SError{http.StatusUnauthorized, nil, err}, w, r)
		return false
	}
	req, err := http.NewRequest("GET", *c.Chey+URL+URL_WS, nil)
	if err != nil {
		LogHTTP(&utils.SError{http.StatusInternalServerError, nil, err}, w, r)
		return false
	}
	req.AddCookie(cookie)
	log.Printf("%#v\n", cookie)
	res, err := http.DefaultTransport.RoundTrip(req) // Avoid redirection
	if err != nil {
		LogHTTP(&utils.SError{http.StatusInternalServerError, nil, err}, w, r)
		return false
	}
	switch res.StatusCode {
	case 200:
		CheyenneLogHTTP(r.URL.String(), res.StatusCode)
		session := &utils.Session{}
		err := json.NewDecoder(res.Body).Decode(session)
		if err != nil {
			LogHTTP(&utils.SError{http.StatusInternalServerError, nil, err}, w, r)
			return false
		}
		res.Body.Close()
		if err := utils.StoreInCookies(c.Store, session, w, r); err != nil {
			LogHTTP(&utils.SError{http.StatusInternalServerError, nil, err}, w, r)
			return false
		}
	case 302:
		url, _ := res.Location()
		CheyenneErrorHTTP(w, url.String(), res.StatusCode)
		return false
	default:
		CheyenneErrorHTTP(w, "", res.StatusCode)
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
		LogHTTP(&utils.SError{http.StatusUnauthorized, nil, err}, w, r)
		return
	}
	if s.Role < 3 {
		LogHTTP(&utils.SError{http.StatusUnauthorized, nil, err}, w, r)
		return
	}
	ah.c.Session = s
	e := ah.fn(ah.c, w, r)
	LogHTTP(e, w, r)
}
