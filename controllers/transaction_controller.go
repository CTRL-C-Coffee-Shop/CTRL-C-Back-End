package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//order harap dipindah ke paling bawah
func GetCart(c *gin.Context) {
	db := connect()

	UserID, _ := strconv.Atoi(c.PostForm("UserID"))
	var cart Cart

	result := db.Where("id_user = ?", UserID).First(&cart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Success fetching product data", "cart": cart})
}

func PostCart(c *gin.Context) { //produk masuk ke cart satu per satu
	UserID, _ := strconv.Atoi(c.PostForm("UserID"))
	ProdID, _ := strconv.Atoi(c.PostForm("ProdID"))
	Amount, _ := strconv.Atoi(c.PostForm("Amount"))

	db := connect()

	newCart := Cart{
		UserID:    UserID,
		ProductID: ProdID,
		Amount:    Amount,
	}

	result := db.Create(&newCart)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to enter product into user's cart"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product successfully append into cart"})
}

func DeleteCart(c *gin.Context) { //delete the cart from specific user and specific product
	db := connect()

	UserID, _ := strconv.Atoi(c.PostForm("UserID"))
	ProdID, _ := strconv.Atoi(c.PostForm("ProdID"))
	var cart Cart

	result := db.Where("id_user = ? AND id_product = ?", UserID, ProdID).Delete(&cart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cart item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product successfully deleted from cart"})
}

func UpdateCart(c *gin.Context) { //only change the amount of product that added into user's cart
	UserID, _ := strconv.Atoi(c.PostForm("UserID"))
	ProdID, _ := strconv.Atoi(c.PostForm("ProdID"))
	Amount, _ := strconv.Atoi(c.PostForm("Amount"))

	cart := Cart {
		UserID:    UserID,
		ProductID: ProdID,
		Amount:    Amount,
	}

	db := connect()
	result := db.Where("id_user = ?", UserID).First(&cart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart item"})
		return
	}
	result2 := db.Save(&cart)
	if result2.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change cart data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cart successfully updated"})
}

//order harap dipindah kebawah sini