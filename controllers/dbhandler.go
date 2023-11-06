package controllers

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ctrl_c+db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil{
		log.Fatal(err)
	}
	return db
}