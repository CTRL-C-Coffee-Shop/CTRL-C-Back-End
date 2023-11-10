package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getKedai(c * gin.Context) {
	db := connect()

	var Kedai []Kedai

	result := db.Table("product").Find(&Kedai) // Perhatikan penggunaan db.Table("product") untuk merujuk tabel "product"

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil produk"})
		return
	}

	c.JSON(http.StatusOK, Kedai)
}