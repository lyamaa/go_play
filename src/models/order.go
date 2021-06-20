package models

type Order struct {
	Model
	TransactionId string      `json:"transaction_id" gorm:"null"`
	UserId        uint        `json:"user_id"`
	Code          string      `json:"code"`
	VendorEmail   string      `json:"vendor_email"`
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	Email         string      `json:"email"`
	Address       string      `json:"address"`
	City          string      `json:"city"`
	Country       string      `json:"country"`
	Zip           string      `json:"zip"`
	Complete      bool        `json:"complete" gorm:"default:false"`
	OrderItem     []OrderItem `json:"order_item" gorm:"foreignKey:OrderId"`
}

type OrderItem struct {
	Model
	OrderId       uint    `json:"order_id"`
	ProductTitle  string  `json:"product_title"`
	Price         float64 `json:"price"`
	Quantity      uint    `json:"quantity"`
	AdminRevenue  float64 `json:"admin_revenue"`
	VendorRevenue float64 `json:"vendor_revenue"`
}
