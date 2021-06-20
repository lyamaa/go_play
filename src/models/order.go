package models

type Order struct {
	Model
	TransactionId string      `json:"transaction_id" gorm:"null"`
	UserId        uint        `json:"user_id"`
	Code          string      `json:"code"`
	VendorEmail   string      `json:"vendor_email"`
	FirstName     string      `json:"-"`
	LastName      string      `json:"-"`
	Name          string      `json:"Name" gorm:"-"`
	Email         string      `json:"email"`
	Address       string      `json:"address"`
	City          string      `json:"city"`
	Country       string      `json:"country"`
	Zip           string      `json:"zip"`
	Complete      bool        `json:"-" gorm:"default:false"`
	Total         float64     `json:"total" gorm:"-"`
	OrderItems    []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
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

func (order *Order) FullName() string {
	return order.FirstName + " " + order.LastName
}

func (order *Order) GetTotal() float64 {
	var total float64 = 0

	for _, orderItem := range order.OrderItems {
		total += orderItem.Price * float64(orderItem.Quantity)
	}

	return total
}
