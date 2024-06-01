package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func Connect() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sensordb")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDB() *sql.DB {
	return db
}
