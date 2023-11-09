package controllers

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()
	 
	email := r.Form.Get("email")
	//get password the change into sha256
	password := r.Form.Get("password")
	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)

	row := db.QueryRow("SELECT * FROM users WHERE email=? AND password =?", email, bs )

	var user User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.UserType); err != nil{
		fmt.Println(err)
		sendErrorResponse(w, "error username or password missmatch")
	} else {
		generateToken(w, user.ID, user.Name, user.Email, user.UserType)
	}
}

func Logout(w http.ResponseWriter, r *http.Request){
	resetUserToken(w)

	var response UserResponse
	response.Status = 200
	response.Message = "Successfully logout"

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