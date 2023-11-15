package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	// Menginisialisasi koneksi database
	db := connect() // Anda perlu mengganti "connect" dengan cara Anda menginisialisasi koneksi database

	// Menggunakan GORM untuk mengambil semua produk dari basis data
	var products []Product // Menggunakan []Product untuk merepresentasikan tabel "product"

	result := db.Table("product").Find(&products) // Perhatikan penggunaan db.Table("product") untuk merujuk tabel "product"

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, products)
}
//order pake mux
func PostOrder(w http.ResponseWriter, r *http.Request){
	db := connectMux()

	UID,_ := strconv.Atoi(r.Form.Get("id_user"))
	ShopID,_ := strconv.Atoi(r.Form.Get("id_kedai"))
	VoucherID,_ := strconv.Atoi(r.Form.Get("id_voucher"))
	total,_ := strconv.Atoi(r.Form.Get("total"))
	var stat = "ongoing"

	//get multiple order, later convert into integers
	str1,_ := r.Form["productID"]
	str2,_ := r.Form["amount"]

	res, errQuery := db.Exec("INSERT INTO `orders`(`id_user`, `id_kedai`, `total_price`, `status`, `id_voucher`) VALUES ('?','?','?','?','?')", UID, ShopID, total, stat, VoucherID)
	if errQuery != nil {
		sendErrorResponse(w, 400, "Failed while inserting transaction")
		return
	}

	order_id,_ := res.LastInsertId() //get id order
	for i := 0; i < len(str1); i++ {
		//insert ke detail transaksi
		_, errQuery := db.Exec("INSERT INTO `detail_orders`(`id_order`, `id_product`, `number`) VALUES ('?','?','?')",
			order_id,
			str1[i],
			str2[i],
		)
		if errQuery != nil {
			fmt.Println(errQuery)
			sendErrorResponse(w, 400, "Failed to record transaction")
			return
		}
	}
	return
}

func sendErrorResponse(w http.ResponseWriter,status int, message string) {
	w.Header().Set("Content=Type", "application/json")
	w.WriteHeader(status)

	errorResponse := ErrorResponse{
		Status:  status,
		Message: message,
	}

	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		log.Println(err)
	}
}