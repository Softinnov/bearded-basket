package models

import (
	"testing"

	"github.com/Softinnov/bearded-basket/server/database"
)

func TestGetUser(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	u, err := GetUser(c, 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if u == nil {
		t.Errorf("Expected a user, got nil")
	}
	su := User{
		Id:       1,
		Pdv:      0,
		Prenom:   "(super)",
		Nom:      "administrateur",
		Role:     9,
		Password: "",
		Login:    "admin",
		FaitPar:  0,
	}
	if su != *u {
		t.Errorf("Expected %#v, go %#v", su, *u)
	}
}

func TestGetUserFromBadId(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	u, err := GetUser(c, 0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if u != nil {
		t.Errorf("Unexpected user, got %#v", u)
	}
}

func TestGetCurrentUser(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	s := &Session{
		Id: 1,
	}
	u, err := GetCurrentUser(c, s)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if u == nil {
		t.Errorf("Expected a user, got nil")
	}
	su := User{
		Id:       1,
		Pdv:      0,
		Prenom:   "(super)",
		Nom:      "administrateur",
		Role:     9,
		Password: "",
		Login:    "admin",
		FaitPar:  0,
	}
	if su != *u {
		t.Errorf("Expected %#v, go %#v", su, *u)
	}
}

func TestGetBadCurrentUser(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	s := &Session{
		Id: 0,
	}
	u, err := GetCurrentUser(c, s)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if u != nil {
		t.Errorf("Unexpected user, got %#v", u)
	}
}

func TestGetCurrentUserNilSession(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	u, err := GetCurrentUser(c, nil)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if u != nil {
		t.Errorf("Unexpected user, got %#v", u)
	}
}

func TestCreateUser(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	u := &User{
		Pdv:      0,
		Nom:      "NomTest",
		Prenom:   "PrenomTest",
		Role:     8,
		Password: "coucou",
		Login:    "loginTest",
		FaitPar:  4,
	}
	s := &Session{
		Id: 1,
	}
	err := CreateUser(c, u, s)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
}
