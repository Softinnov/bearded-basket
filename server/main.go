package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/handlers"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/sessions"
)

var (
	dbuspw  = flag.String("dbuspw", "root:", "database, usage: user:passwd@addr/dbname")
	dbname  = flag.String("dbname", "prod", "database, usage: user:passwd@addr/dbname")
	conf    = flag.String("conf", "./config.json", "configuration file")
	encrypt = []byte("123456789")
)

func fatal(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func notfatal(s string) {
	log.Printf(s)
}

func main() {

	var httpDB database.Db

	flag.Parse()

	up := strings.Split(*dbuspw, ":")
	if len(up) != 2 {
		log.Fatal("Invalid user:password format")
	}

	context := &utils.Context{
		Store:   sessions.NewCookieStore(encrypt),
		HTTPdb:  &httpDB,
		Session: nil,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP)

	go func(s <-chan os.Signal, hdb *database.Db) {
		for {
			f, e := os.Open(*conf)
			fatal(e)
			d := struct {
				Client []string `json:"client"`
				HttpDB []string `json:"httpdb"`
			}{}
			e = json.NewDecoder(f).Decode(&d)
			fatal(e)
			log.Printf("%+v", d)

			if len(d.HttpDB) == 0 {
				notfatal("no httpdb found")
				time.Sleep(5 * time.Second)
				continue
			}

			*hdb = database.Db{
				Host:     d.HttpDB[0],
				User:     up[0],
				Password: up[1],
				Database: *dbname,
			}
			log.Printf("New database configuration %s\n", hdb.Host)

			if len(d.Client) == 0 {
				notfatal("no client found")
				time.Sleep(5 * time.Second)
				continue
			}
			context.Client = d.Client[0]
			log.Printf("New client configuration %s\n", context.Client)

			// Block until a signal is received.
			ss := <-s
			log.Println("Got signal: ", ss)
		}
	}(c, &httpDB)

	handlers.Init(context)

	http.ListenAndServe(":8002", nil)
}
