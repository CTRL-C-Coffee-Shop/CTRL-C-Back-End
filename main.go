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
	r.GET("/userlogout", controllers.Logout)

	// Your existing /product route
	r.GET("/product", controllers.Authenticate(false), controllers.GetAllProducts)

	r.Run("localhost:8080")
}
