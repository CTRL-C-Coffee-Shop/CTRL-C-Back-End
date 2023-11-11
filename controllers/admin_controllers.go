package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrder(c *gin.Context) {
	db := connect()

	var orders []Order

	result := db.Preload("Orders").Find(&orders) // Preload orders details

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil produk"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func UpdateOrders(c *gin.Context) {
	db := connect()

	orderID := c.PostForm("OrderID")
	stat := c.PostForm("Status")

	if orderID == "" || stat == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Harap isi semua field"})
		return
	}

	//get data
	var order Order
	if err := db.First(&order, orderID).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Order not found"})
		return
	}

	//update
	order.status = stat
	if err := db.Save(&order).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Gagal Update status"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order Updated"})
}
