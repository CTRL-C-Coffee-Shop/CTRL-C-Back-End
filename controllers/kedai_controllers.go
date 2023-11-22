package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKedai(c *gin.Context) {
	db := connect()

	var Kedai []Kedai

	result := db.Table("kedai").Find(&Kedai) // Perhatikan penggunaan db.Table("product") untuk merujuk tabel "product"

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data successfully retrieved", "Kedai": Kedai})
}
