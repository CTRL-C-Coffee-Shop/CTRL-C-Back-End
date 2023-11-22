package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllVouchers(c *gin.Context) {
	db := connect()

	// Menggunakan GORM untuk mengambil semua produk dari basis data
	var vouchers []Voucher

	result := db.Table("voucher").Find(&vouchers)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve voucher"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully retrieved vouchers", "vouchers": vouchers})
}
