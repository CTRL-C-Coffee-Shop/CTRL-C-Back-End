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

type Order struct {
	IDOrder   uint   `gorm:"column:id_order;primaryKey" json:"id_order"`
	UserID    uint   `gorm:"column:id_user;not null" json:"user_id"`
	ShopID    uint   `gorm:"column:id_kedai;not null" json:"id_kedai"`
	VoucherID uint   `gorm:"column:id_voucher;not null" json:"voucher_id"`
	Date      string `gorm:"column:date;not null" json:"date"`
	Price     uint   `gorm:"column:price;not null" json:"price"`
	Status    string `gorm:"column:status;not null" json:"status"`
}
type Voucher struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Discount    uint   `json:"discount"`
	Number      uint   `json:"number"`
}
