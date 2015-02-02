package utils

import (
	"net/http"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/gorilla/sessions"
	"github.com/mitchellh/mapstructure"
)

type Context struct {
	Store   *sessions.CookieStore
	HTTPdb  *database.Db
	Chey    *string
	Session *Session
}

type Session struct {
	Id    int64    `json:"id"`
	Name  []string `json:"name"`
	Role  int8     `json:"role"`
	PdvId int      `json:"pdvid"`
}

func StoreInCookies(store *sessions.CookieStore, s *Session, w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "session-go")
	if err != nil {
		return err
	}

	session.Values["id"] = s.Id
	session.Values["name"] = s.Name
	session.Values["role"] = s.Role
	session.Values["pdvid"] = s.PdvId

	err = session.Save(r, w)
	if err != nil {
		return err
	}
	return nil
}

func SessionFromCookies(store *sessions.CookieStore, r *http.Request) (*Session, error) {
	var s Session

	se, e := store.Get(r, "session-go")
	if e != nil {
		return nil, e
	}
	e = mapstructure.Decode(se.Values, &s)
	if e != nil {
		return nil, e
	}
	return &s, nil
}
