package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Model
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Email     string   `json:"email" gorm:"unique"`
	Password  []byte   `json:"-"`
	IsVendor  bool     `json:"-"`
	Revenue   *float64 `json:"revenue,omitempty" gorm:"-"`
}

// Password HASHING
func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	user.Password = hashedPassword
}

// COMPARE PASSWORD
func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

type Admin User //aliases
func (admin *Admin) CalculateRevenue(db *gorm.DB) {
	var orders []Order
	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   admin.Id,
		Complete: true,
	})
	var revenue float64 = 0
	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.AdminRevenue
		}
	}
	admin.Revenue = &revenue
}

type Vendor User

func (vendor *Vendor) CalculateRevenue(db *gorm.DB) {
	var orders []Order
	db.Preload("OrderItems").Find(&orders, &Order{
		UserId:   vendor.Id,
		Complete: true,
	})
	var revenue float64 = 0
	for _, order := range orders {
		for _, orderItem := range order.OrderItems {
			revenue += orderItem.VendorRevenue
		}
	}
	vendor.Revenue = &revenue
}
