package models

import (
	"testing"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/sessions"
)

func newTestContext(t *testing.T) *utils.Context {
	db := database.Open("admin:admin@(db_test:3306)/prod_test")

	chey := ""
	c := &utils.Context{
		Store: sessions.NewCookieStore([]byte("123456789")),
		DB:    db,
		Chey:  &chey,
	}

	return c
}

func TestGetRoles(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	rs, err := GetRoles(c)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if len(rs) != 9 {
		t.Errorf("Expected 9 roles, got %d", len(rs))
	}
}

func TestGetRole(t *testing.T) {
	c := newTestContext(t)
	defer database.Close(c.DB)

	r, err := GetRole(c, 1)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	sr := Role{
		Id:      1,
		Libelle: "hôte de caisse",
	}
	if r == nil {
		t.Errorf("Expected a role, got nil")
	}
	if sr != *r {
		t.Errorf("Expected %#v, got %#v", sr, *r)
	}
}
