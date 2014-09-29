package main

import (
	"flag"
	"net/http"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/handlers"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/sessions"
)

var dbf = flag.String("db", "root:@/prod", "database, usage: user:password@addr/dbname")

func main() {
	flag.Parse()

	db := database.Open(*dbf)
	defer database.Close(db)

	context := &utils.Context{
		Store: sessions.NewCookieStore([]byte("123456789")),
		DB:    db,
	}

	handlers.Init(context)

	http.ListenAndServe(":8002", nil)
}
