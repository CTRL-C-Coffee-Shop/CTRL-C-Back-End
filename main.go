package main

import (
	"CTRL-C-Back-End/controllers"
	"net/http"

	// Import the CORS middleware
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Use CORS middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	})

	// Your existing routes
	r.POST("/register", controllers.Register)
	r.POST("/userlogin", controllers.Login)
	r.POST("/getorder", controllers.Authenticate(false), controllers.GetOrder)
	r.POST("/createorder", controllers.Authenticate(false), controllers.CreateOrder)
	r.POST("/getcart", controllers.Authenticate(false), controllers.GetCart)
	r.POST("/postcart", controllers.Authenticate(false), controllers.PostCart)
	r.POST("/deletecart", controllers.Authenticate(false), controllers.DeleteCart)
	r.POST("/deleteallcart", controllers.Authenticate(false), controllers.DeleteAllCart)
	//admin order
	r.GET("/getorderadmin", controllers.Authenticate(true), controllers.GetOrderAdmin)
	r.POST("/updateorderstatus", controllers.Authenticate(true), controllers.UpdateOrderStatus)

	// voucher, kedai and product
	r.GET("/getvoucher", controllers.Authenticate(false), controllers.GetAllVouchers)
	r.GET("/product", controllers.Authenticate(false), controllers.GetAllProducts)
	r.GET("/getkedai", controllers.Authenticate(false), controllers.GetKedai)

	r.Run("localhost:8080")
}
