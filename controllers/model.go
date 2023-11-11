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

type OrderDetail struct {
	OrderID   uint `gorm:"primaryKey"`
	ProductID uint `gorm:"primaryKey"`
	Amount    uint `gorm:"not null"`
}

type Order struct {
	OrderID   uint          `gorm:"primaryKey"`
	UserID    uint          `gorm:"not null"`
	ShopID    uint          `gorm:"not null"`
	VoucherID uint          `gorm:"not null"`
	orders    []OrderDetail `gorm:"not null"`
	date      string        `gorm:"not null"`
	price     uint          `gorm:"not null"`
	status    string        `gorm:"not null"`

type Voucher struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Discount    uint   `json:"price"`
	Number      uint   `json:"type"`
}