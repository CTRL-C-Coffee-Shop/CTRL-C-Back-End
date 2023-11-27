package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrderAdmin(c *gin.Context) {
	db := connect()

	var orders []OrderAdmin

	// Use Preload to include the associated Voucher data
	result := db.Table("orders").Find(&orders)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	// Iterate through orders and split the date and time
	for i := range orders {
		datetime, err := time.Parse(time.RFC3339, orders[i].Date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse datetime"})
			return
		}

		// Set the separated date and time values
		orders[i].DateOnly = datetime.Format("2006-01-02") // Format "YYYY-MM-DD"
		orders[i].TimeOnly = datetime.Format("15:04:05")   // Format "HH:MM:SS"
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orders retrieved successfully", "orders": orders})
}

func UpdateOrderStatus(c *gin.Context) {
	db := connect()

	// Get order ID from request parameters
	orderID := c.PostForm("Order_Id")

	// Find the order by ID
	var order OrderAdmin
	result := db.Table("orders").First(&order, orderID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Get the new status from request form data
	newStatus := c.PostForm("Status")

	// Check if the new status is empty
	if newStatus == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status cannot be empty"})
		return
	}

	// Update the status field
	result = db.Model(&order).Update("status", newStatus)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully", "order": order})
}
