package controllers

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

var db *gorm.DB

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
    var products []Product

    // Retrieve all products from the database
    db.Find(&products)

    // Serialize the products to JSON and send as the response
    json.NewEncoder(w).Encode(products)
}
