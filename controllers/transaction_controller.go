package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// order harap dipindah ke paling bawah
func GetCart(c *gin.Context) {
	db := connect()

	UserID, _ := strconv.Atoi(c.PostForm("UserID"))

	// Menggunakan slice untuk menyimpan semua produk dalam keranjang
	var cart []Cart

	result := db.Where("id_user = ?", UserID).Find(&cart)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart items"})
		return
	}

	// Menyiapkan slice untuk menyimpan data item keranjang yang akan dikembalikan
	var cartData []gin.H

	// Loop melalui setiap item dalam keranjang
	for _, item := range cart {
		// Mengambil data produk terkait dari tabel Product berdasarkan ProductID
		var product Product
		db.First(&product, item.ProductID)

		// Menambahkan data item dan data produk ke dalam slice cartData
		cartData = append(cartData, gin.H{
			"UserID":    item.UserID,
			"ProductID": item.ProductID,
			"Amount":    item.Amount,
			"Product":   product, // Menambahkan data produk ke dalam respons
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success fetching cart data", "cart": cartData})
}

func PostCart(c *gin.Context) {
	UserID, _ := strconv.Atoi(c.PostForm("UserID"))
	ProdID, _ := strconv.Atoi(c.PostForm("ProdID"))
	Amount, _ := strconv.Atoi(c.PostForm("Amount"))
	Warmnth, _ := strconv.Atoi(c.PostForm("Warmnth"))
	Size, _ := strconv.Atoi(c.PostForm("Size"))
	SugarLvl, _ := strconv.Atoi(c.PostForm("SugarLvl"))

	db := connect()

	// Cari cart dengan user id dan product id yang diberikan
	var cart Cart
	result := db.Where("id_user = ? AND id_product = ?", UserID, ProdID).First(&cart)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Jika record tidak ditemukan, tambahkan sebagai entri baru
			newCart := Cart{
				UserID:    UserID,
				ProductID: ProdID,
				Amount:    Amount,
				Warmth:    Warmnth,
				Size:      Size,
				SugarLvl:  SugarLvl,
			}
			result := db.Create(&newCart)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add new cart item"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart item"})
			return
		}
	} else {
		// Update jumlah dengan menambahkan nilai yang baru ke nilai yang sudah ada
		newAmount := cart.Amount + Amount

		// Update semua field yang perlu diupdate
		result := db.Model(&cart).Updates(Cart{
			Amount: newAmount,
		})
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update cart data"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart successfully updated"})
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

// order harap dipindah kebawah sini
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
	currentDate := time.Now().Format("2006-01-02 15:04:05")

	// Create a new order
	newOrder := OrderUser{
		UserID:    uint(userID),
		ShopID:    uint(shopID),
		VoucherID: uint(voucherID),
		Price:     uint(totalPrice),
		Date:      currentDate,
		Status:    status,
	}

	if err := db.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Iterate through products and create order details
	productIDs := c.PostFormArray("productID")
	amounts := c.PostFormArray("amount")
	warmths := c.PostFormArray("warmth")
	sizes := c.PostFormArray("size")
	sugarLvls := c.PostFormArray("sugarLvl")

	// Assuming newOrder is already defined and available
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

		warmth, err := strconv.Atoi(warmths[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid warmth"})
			return
		}

		size, err := strconv.Atoi(sizes[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid size"})
			return
		}

		sugarLvl, err := strconv.Atoi(sugarLvls[i])
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sugar level"})
			return
		}

		orderDetail := OrderDetail{
			IDOrder:   newOrder.IDOrder,
			ProductID: uint(productID),
			Amount:    uint(amount),
			Warmth:    warmth,
			Size:      size,
			SugarLvl:  sugarLvl,
		}

		if err := db.Create(&orderDetail).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order detail"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}
