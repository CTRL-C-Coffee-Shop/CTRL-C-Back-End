package main

import (
	"CTRL-C-Back-End/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// user
	r.POST("/register", controllers.Register)
	r.GET("/userlogin", controllers.Login)
	r.GET("/userlogout", controllers.Logout)

	//product
	r.GET("/product", controllers.Authenticate(false), controllers.GetAllProducts)

	r.Run("localhost:8080")
}
