package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)

	if err != nil {
		log.Panic(err)
	}
	if db == nil {
		log.Panic("db nil")
	}
}
