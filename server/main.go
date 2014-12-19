package main

import (
	"flag"
	"net/http"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/handlers"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var dbf = flag.String("db", "root:@/prod", "database, usage: user:password@addr/dbname")
var cheyf = flag.String("chey", "http://localhost:8002", "cheyenne, usage: http://host:port")

func main() {
	flag.Parse()

	db := database.Open(*dbf)
	defer database.Close(db)

	context := &utils.Context{
		Store:   sessions.NewCookieStore(securecookie.GenerateRandomKey(32)),
		DB:      db,
		Chey:    cheyf,
		Session: nil,
	}

	handlers.Init(context)

	http.ListenAndServe(":8002", nil)
}
