package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
func CreateOrder(c *gin.Context) {
	db := connect()
	userIDStr := c.PostForm("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	shopIDStr := c.PostForm("id_kedai")
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop ID"})
		return
	}

	voucherIDStr := c.PostForm("id_voucher")
	voucherID, err := strconv.Atoi(voucherIDStr)
	if err != nil {
		// Handle if voucher ID is optional or not present
		voucherID = 0
	}

	totalPriceStr := c.PostForm("total")
	totalPrice, err := strconv.Atoi(totalPriceStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid total price"})
		return
	}

	var status = "ongoing"

	// Create a new order
	newOrder := OrderUser{
		UserID:    uint(userID),
		ShopID:    uint(shopID),
		VoucherID: uint(voucherID),
		Price:     uint(totalPrice),
		Status:    status,
	}

	if err := db.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Iterate through products and create order details
	productIDs := c.PostFormArray("productID")
	amounts := c.PostFormArray("amount")

	fmt.Println(amounts)

	for i := 0; i < len(productIDs); i++ {
		productID, err := strconv.Atoi(productIDs[i])

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
			return
		}

		amount, err := strconv.Atoi(amounts[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
			return
		}

		orderDetail := OrderDetail{
			IDOrder:   newOrder.IDOrder,
			ProductID: uint(productID),
			Amount:    uint(amount),
		}

		if err := db.Create(&orderDetail).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order detail"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
