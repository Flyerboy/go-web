package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func getDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/video")
	if err != nil {
		panic(err.Error())
	}
	return db
}