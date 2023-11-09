package controllers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:abc@tcp(127.0.0.1:3306)/ctrl_c_db?charset=utf8mb4&parseTime=True&loc=Local"
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
