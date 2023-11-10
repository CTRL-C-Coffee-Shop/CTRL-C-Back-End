package controllers

type User struct {
	ID       uint   `gorm:"primaryKey"`
	FullName string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	AccType  bool   `gorm:"not null"`
}
type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Type        string  `json:"type"`
	URL         string  `json:"url"`
}

type Kedai struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}