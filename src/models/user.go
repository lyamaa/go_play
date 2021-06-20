package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	IsVendor  bool   `json:"-"`
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
