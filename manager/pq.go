package manager

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
var dbConn *sql.DB

func connPG() {
	connStr := "postgres://postgres:postgres@127.0.0.1:5432/yt"
	var err error
	dbConn, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func joinGroupQurey(uid uint64) bool {
	return true
}
