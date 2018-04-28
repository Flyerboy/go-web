package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init()  {
	db, err := sql.Open("mysql", "root:root@/video")
	if err != nil {
		panic(err.Error())
	}
	DB = db
}