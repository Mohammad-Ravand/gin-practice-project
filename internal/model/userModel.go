// user.go
package models

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Email    string
	Password string
}

func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
func GetUserByEmail(db *gorm.DB, email string, user *User) error {
	return db.Where("email", email).First(&user).Error // find product with code D42
}
