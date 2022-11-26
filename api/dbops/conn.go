package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:12345678@(localhost:3306)/videoserver?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
}
