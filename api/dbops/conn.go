package dbops

import (
	"database/sql"

	// Avoid blank import in non-main
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:c11090201@/video_server?charset=utf8mb4")
	if err != nil {
		// terminate app
		panic(err.Error())
	}
}
