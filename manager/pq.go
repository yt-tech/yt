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
func loginQurey(uid uint64) bool {
	// mlog.Printf("select data %d\n", uid)
	// row := dbConn.QueryRow("SELECT uid FROM yt_user WHERE uid = $1", uid)
	// var i uint64
	// err := row.Scan(&i)
	// if err != nil {
	// 	mlog.Println(err)
	// 	return false
	// }
	return true
}
func joinGroupQurey(uid uint64) bool {
	return true
}
