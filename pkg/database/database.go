package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() *sql.DB {
	var err error
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	connection := fmt.Sprintf("root:root@tcp(%s:3306)/sensordb", dbHost)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetDB() *sql.DB {
	return db
}
