package dbops

import (
	"database/sql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root@tcp(localhost:3306)/video_server?charset=utf8")
	//dbConn, err = sql.Open("mysql", "root@(127.0.0.1:3306)/video_server?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
}
