package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrder(c *gin.Context) {
	db := connect()

	var orders []OrderUser
	userId := c.PostForm("id")
	result := db.Table("orders").
		Where("id_user = ?", userId).
		Preload("OrderDetail").
		Preload("OrderDetail.Product").
		Preload("Voucher").
		Preload("Kedai").
		Find(&orders)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully retrieved orders", "orders": orders})
}
