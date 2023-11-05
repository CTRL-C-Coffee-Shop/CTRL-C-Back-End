package controllers

type User struct {
	ID       int    `json: "id"`
	Name     string `json: "full_name"`
	Email    int    `json: "email"`
	UserType bool   `json: "acc_type"`
}