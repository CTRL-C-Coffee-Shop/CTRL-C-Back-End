package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	// Menginisialisasi koneksi database
	db := connect() // Anda perlu mengganti "connect" dengan cara Anda menginisialisasi koneksi database

	// Menggunakan GORM untuk mengambil semua produk dari basis data
	var products []Product // Menggunakan []Product untuk merepresentasikan tabel "product"

	result := db.Table("product").Find(&products) // Perhatikan penggunaan db.Table("product") untuk merujuk tabel "product"

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, products)
}
