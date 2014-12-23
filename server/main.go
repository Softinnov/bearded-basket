package main

import (
	"flag"
	"net/http"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/handlers"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/sessions"
)

var (
	dbf     = flag.String("db", "root:@/prod", "database, usage: user:password@addr/dbname")
	cheyf   = flag.String("chey", "http://localhost:8002", "cheyenne, usage: http://host:port")
	encrypt = []byte("123456789")
)

func main() {
	flag.Parse()

	db := database.Open(*dbf)
	defer database.Close(db)

	context := &utils.Context{
		Store:   sessions.NewCookieStore(encrypt),
		DB:      db,
		Chey:    cheyf,
		Session: nil,
	}

	handlers.Init(context)

	http.ListenAndServe(":8002", nil)
}
