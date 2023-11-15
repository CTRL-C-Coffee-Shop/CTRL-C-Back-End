package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:abc@tcp(127.0.0.1:3306)/ctrl_c_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Tanpa Password
	// dsn := "root:@tcp(127.0.0.1:3306)/ctrl_c+db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	// Uji koneksi ke database
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to get database instance")
	}
	// Pastikan untuk menutup koneksi database ketika selesai
	err = sqlDB.Ping()
	if err != nil {
		panic("Failed to ping database: " + err.Error())
	}
	return db
}

func connectMux() *sql.DB {
	_ = godotenv.Load()
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")
	dbaddress := fmt.Sprintf("root:@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_port, db_name)
	db, err := sql.Open("mysql", dbaddress)
	if err != nil {
		log.Fatal()
	}
	return db
}
