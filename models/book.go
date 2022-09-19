package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title      string
	Author     string
	OwnerId    uint
	Borrowed   bool
	StartDate  time.Time
	EndDate    time.Time
	BorrowerId uint
	Waitlist   []User `gorm:"many2many:waitlist"`
}
