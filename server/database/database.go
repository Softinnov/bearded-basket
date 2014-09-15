package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(addr string) {
	var err error

	DB, err = sql.Open("mysql", addr+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	DB.Close()
}
