package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	Host     string
	User     string
	Password string
	Database string
}

type QueryResult struct {
	Columns []string         `json:"columns,omitempty"`
	Data    [][]*string      `json:"data,omitempty"`
	Infos   map[string]int64 `json:"infos,omitempty"`
	Error   string           `json:"error,omitempty"`
}

func (db *Db) fetch(query string) (*QueryResult, error) {

	c := http.Client{}

	req, e := http.NewRequest("POST", query, nil)
	if e != nil {
		return nil, e
	}
	req.SetBasicAuth(db.User, db.Password)

	log.Printf("%s\n", query)
	r, e := c.Do(req)
	if e != nil {
		return nil, e
	}
	defer r.Body.Close()

	dbq := &QueryResult{}

	e = json.NewDecoder(r.Body).Decode(dbq)
	if e != nil {
		return nil, e
	}
	if dbq.Error != "" {
		return nil, fmt.Errorf(dbq.Error)
	}

	return dbq, nil
}

// Response of type:
//
// {
//   "columns": ["id", "name", "age"]
//   "data": [
//             ["0", "john", "18"],
//             ["1", null, "42"]
//         ],
//   "error": null
// }
func (db *Db) Query(query string) (*QueryResult, error) {
	q := "http://" + db.Host +
		"/query/" + db.Database + "/" + query

	return db.fetch(q)
}

// Response of type:
//
// {
//   "infos": {
//           "lastInsertId": 43,
//           "rowsAffected": 1
//         },
//   "error": null
// }
func (db *Db) Exec(query string) (*QueryResult, error) {
	q := "http://" + db.Host +
		"/exec/" + db.Database + "/" + query

	return db.fetch(q)
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
