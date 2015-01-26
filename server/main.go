package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Softinnov/bearded-basket/server/database"
	"github.com/Softinnov/bearded-basket/server/handlers"
	"github.com/Softinnov/bearded-basket/server/utils"
	"github.com/gorilla/sessions"
	capi "github.com/hashicorp/consul/api"
)

var (
	dbhost  = flag.String("dbhost", "", "location of sql database")
	dbuspw  = flag.String("dbuspw", "root:", "database, usage: user:passwd@addr/dbname")
	dbname  = flag.String("dbname", "prod", "database, usage: user:passwd@addr/dbname")
	cheyf   = flag.String("chey", "http://localhost:8002", "cheyenne, usage: http://host:port")
	encrypt = []byte("123456789")
)

func getAddrFromConsul(a string) (string, string, error) {
	c := capi.DefaultConfig()
	c.Address = "consul:8500"
	cl, e := capi.NewClient(c)
	if e != nil {
		return "", "", e
	}
	cas, _, e := cl.Catalog().Service(a, "", nil)
	if e != nil {
		return "", "", e
	}
	for _, ca := range cas {
		if ca.ServiceName == a {
			log.Printf("Found %s:%s\n", ca.Address, strconv.Itoa(ca.ServicePort))
			return ca.Address, strconv.Itoa(ca.ServicePort), nil
		}
	}
	return "", "", fmt.Errorf("Nothing found")
}

func main() {
	flag.Parse()

	a, p, e := getAddrFromConsul("db")
	if e != nil {
		log.Fatal(e)
	}
	db := database.Open(*dbuspw + "@(" + a + ":" + p + ")/" + *dbname)
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
