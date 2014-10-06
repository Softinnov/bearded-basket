package utils

import (
	"database/sql"

	"github.com/gorilla/sessions"
)

type Context struct {
	Store *sessions.CookieStore
	DB    *sql.DB
	Chey  *string
}
