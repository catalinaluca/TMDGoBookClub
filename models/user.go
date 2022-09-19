package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName     string
	LastName      string
	Username      string
	Email         string
	Password      string
	Roles         []Role `gorm:"many2many:user_roles"`
	OwnedBooks    []Book `gorm:"foreignKey:OwnerId"`
	BorrowedBooks []Book `gorm:"foreignKey:BorrowerId"`
	Wishlist      []Book `gorm:"many2many:wishlist"`
}
