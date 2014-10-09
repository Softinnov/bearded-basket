package models

import (
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
)

type Session struct {
	Id    int      `json:"id"`
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

func GetSessionFromCookies(store *sessions.CookieStore, r *http.Request) (*Session, error) {
	session, err := store.Get(r, "session-go")
	if err != nil {
		return nil, err
	}
	id, ok := session.Values["id"].(int)
	if !ok {
		return nil, errors.New("bad session")
	}
	name, ok := session.Values["name"].([]string)
	if !ok {
		return nil, errors.New("bad session")
	}
	role, ok := session.Values["role"].(int8)
	if !ok {
		return nil, errors.New("bad session")
	}
	pdvid, ok := session.Values["pdvid"].(int)
	if !ok {
		return nil, errors.New("bad session")
	}
	s := &Session{
		Id:    id,
		Name:  name,
		Role:  role,
		PdvId: pdvid,
	}
	return s, nil
}
