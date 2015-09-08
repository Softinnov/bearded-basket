package models

import (
	"testing"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/utils"
)

func newTestContext(t *testing.T) *utils.Context {
	db := database.Db{
		Host:     "db-test:6033",
		UsPwd:    "admin:admin",
		Database: "prod_test",
	}

	c := &utils.Context{
		HTTPdb: &db,
	}

	return c
}

func TestGetRoles(t *testing.T) {
	c := newTestContext(t)

	rs, err := GetRoles(c)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if len(rs) != 9 {
		t.Errorf("Expected 9 roles, got %d", len(rs))
	}
}
