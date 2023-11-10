package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllVouchers(c *gin.Context) {
	db := connect() 

	// Menggunakan GORM untuk mengambil semua produk dari basis data
	var vouchers []Product 

	result := db.Table("voucher").Find(&vouchers) 

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil voucher"})
		return
	}

	c.JSON(http.StatusOK, vouchers)
}
