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
	IDOrder   uint    `gorm:"primaryKey;column:id_order" json:"id_order"`
	ProductID uint    `gorm:"primaryKey;column:id_product" json:"id_product"`
	Amount    uint    `gorm:"column:number;not null" json:"amount"`
	Warmth    int     `gorm:"column:warmth" json:"warmth"`
	Size      int     `gorm:"column:size" json:"size"`
	SugarLvl  int     `gorm:"column:sugarLvl" json:"sugarLvl"`
	Product   Product `gorm:"foreignKey:ProductID"`
}

type OrderAdmin struct {
	IDOrder   uint   `gorm:"column:id_order;primaryKey" json:"id_order"`
	UserID    uint   `gorm:"column:id_user;not null" json:"user_id"`
	ShopID    uint   `gorm:"column:id_kedai;not null" json:"id_kedai"`
	VoucherID uint   `gorm:"column:id_voucher;not null" json:"voucher_id"`
	Date      string `gorm:"column:date;not null" json:"date"`
	Price     uint   `gorm:"column:total_price;not null" json:"price"`
	Status    string `gorm:"column:status;not null" json:"status"`
	// Fields for separated date and time
	DateOnly string `json:"date_only"`
	TimeOnly string `json:"time_only"`
}
type OrderUser struct {
	IDOrder     uint          `gorm:"column:id_order;primaryKey" json:"id_order"`
	UserID      uint          `gorm:"column:id_user;not null" json:"user_id"`
	ShopID      uint          `gorm:"column:id_kedai;not null" json:"id_kedai"`
	VoucherID   uint          `gorm:"column:id_voucher;not null" json:"voucher_id"`
	Date        string        `gorm:"column:date;not null" json:"date"`
	Price       uint          `gorm:"column:total_price;not null" json:"price"`
	Status      string        `gorm:"column:status;not null" json:"status"`
	DateOnly    string        `json:"date_only"`
	TimeOnly    string        `json:"time_only"`
	OrderDetail []OrderDetail `gorm:"foreignKey:IDOrder" json:"order_detail"`
	Voucher     Voucher       `gorm:"foreignKey:VoucherID" json:"voucher"`
	Kedai       Kedai         `gorm:"foreignKey:ShopID" json:"kedai"`
}
type Voucher struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Discount    uint   `json:"discount"`
	Number      uint   `json:"number"`
}
type Cart struct {
	UserID    int `gorm:"column:id_user;primaryKey;not null" json:"id_user"`
	ProductID int `gorm:"column:id_product;primaryKey;not null" json:"id_product"`
	Amount    int `gorm:"column:amount;not null" json:"amount"`
	Warmth    int `gorm:"column:warmth" json:"warmth"`
	Size      int `gorm:"column:size" json:"size"`
	SugarLvl  int `gorm:"column:sugarLvl" json:"sugarLvl"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (Kedai) TableName() string {
	return "kedai"
}
func (Product) TableName() string {
	return "product"
}
func (Voucher) TableName() string {
	return "voucher"
}
func (OrderUser) TableName() string {
	return "orders"
}
func (OrderDetail) TableName() string {
	return "detail_orders"
}

func (OrderAdmin) TableName() string {
	return "orders"
}
func (Cart) TableName() string {
	return "cart"
}
