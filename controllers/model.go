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

type OrderDetail struct {
	OrderID   int `gorm:"primaryKey"`
	ProductID int `gorm:"primaryKey"`
	Amount    int `gorm:"not null"`
}

type Order struct {
	OrderID   uint          `gorm:"primaryKey"`
	UserID    uint          `gorm:"not null"`
	ShopID    uint          `gorm:"not null"`
	VoucherID uint          `gorm:"not null"`
	Details   []OrderDetail `gorm:"foreignKey:OrderID"` // Specify the foreign key relationship
	Date      string        `gorm:"not null"`
	Price     uint          `gorm:"not null"`
	Status    string        `gorm:"not null"`
}
type Voucher struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Discount    uint   `json:"price"`
	Number      uint   `json:"type"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}