package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//order harap dipindah ke paling bawah
func GetCart(c *gin.Context) {//get all cart from user x
	db := connect()

	UserID, _ := strconv.Atoi(c.PostForm("UserID"))
	var cart Cart

	result := db.Where("id_user = ?", UserID).Find(&cart)
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
	Warmnth, _ := strconv.Atoi(c.PostForm("Warmnth"))
	Size, _ := strconv.Atoi(c.PostForm("Size"))
	SugarLvl, _ := strconv.Atoi(c.PostForm("SugarLvl"))

	db := connect()

	newCart := Cart{
		UserID:    UserID,
		ProductID: ProdID,
		Amount:    Amount, 
		Warmnth: Warmnth,
		Size: 	Size,
		sugarLvl: SugarLvl,
	}

	result := db.Create(&newCart)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Failed to enter product into user's cart"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product successfully append into cart"})
}

func DeleteCart(c *gin.Context) { //delete product x from user y in cart
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

func DeleteAllCart(c *gin.Context) { //delete all cart from user x
	db := connect()

	UserID, _ := strconv.Atoi(c.PostForm("UserID"))
	var cart Cart

	result := db.Where("id_user = ?", UserID).Delete(&cart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete cart item"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data successfully deleted from cart"})
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