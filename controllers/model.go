package controllers

type User struct {
	ID       int    `json: "id"`
	Name     string `json: "full_name"`
	Email    string `json: "email"`
	UserType bool   `json: "acc_type"`
}

type UserResponse struct {
	Status  int    `json: "id"`
	Message string `json: "name"`
	Data    User   `json:"Data"`
}

type ErrorResponse struct {
	Status  int    `json: "id"`
	Message string `json: "name"`
}

type Product struct {
	ID       int    `json: "id"`
	Name     string `json: "item_name"`
	Desc     string `json: "item_desc"`
	Price    int    `json: "price"`
	MenuType string `json: "item_type"`
	Url      string `json: "url"`
}

type ProductResponse struct {
	Status  int       `json: "id"`
	Message string    `json: "name"`
	Data    []Product `json:"Data"`
}

type Shop struct {
	ID      int    `json: "id"`
	Name    string `json: "shop_name"`
	Address string `json: "address"`
}