package database

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Open(addr string) *sql.DB {
	var err error

	db, err := sql.Open("mysql", addr+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Close(db *sql.DB) {
	db.Close()
}

func BuildSqlSets(b []byte) (string, error) {
	var buf bytes.Buffer
	var d map[string]interface{}

	e := json.Unmarshal(b, &d)
	if e != nil {
		return "", e
	}
	f := false
	for k, v := range d {
		if f {
			buf.WriteString(", ")
		}
		switch v.(type) {
		case string:
			buf.WriteString(k + "=" + fmt.Sprintf("%q", v))
		default:
			buf.WriteString(k + "=" + fmt.Sprintf("%v", v))
		}
		f = true
	}
	return buf.String(), nil
}
