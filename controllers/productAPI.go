package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)


func getProductsHandler(w http.ResponseWriter, r *http.Request) {
    db := connectDB()
	defer db.Close()

	query := "SELECT * FROM `product`"

	rows, err := db.Query(query)
	if err != nil{
		log.Println(err)
		sendErrorResponse(w,"Something went wrong, please try again.")
		return
	}

	var product Product
	var products []Product
	var stat int
	var mess string

	for rows.Next(){
		if err := rows.Scan(&product.ID, &product.Name, &product.Desc, &product.Price, &product.MenuType, &product.Url); err != nil {
			log.Println(err)
			stat = 400
			mess = "Something wrong with request"
			return
		}else{
			stat = 200
			mess = "Successfully return"
			products = append(products, product)
		}
	}

	resp := ProductResponse{
		Status: stat,
		Message: mess,
		Data: products,
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(resp)
}
