package main

import (
	"flag"
	"net/http"

	"github.com/ghigt/ext_users/server/database"
	"github.com/ghigt/ext_users/server/handlers"
	"github.com/ghigt/ext_users/server/utils"
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
