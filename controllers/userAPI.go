package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()
	 
	name := r.URL.Query()

	row := db.QueryRow("SELECT * FROM users WHERE name=?", name)

	var user User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.UserType); err != nil {
		fmt.Println(err)
		sendErrorResponse(w, "error user not found")
	} else {
		generateToken(w, user.ID, user.Name, user.Email, user.UserType)
	}
}

func Logout(w http.ResponseWriter, r *http.Request){
	resetUserToken(w)

	var response UserResponse
	response.Status = 200
	response.Message = "Successfully login"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}


func sendErrorResponse(w http.ResponseWriter, Message string) {
	var response ErrorResponse
	response.Status = 404
	response.Message = Message

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}