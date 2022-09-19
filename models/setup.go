package models

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "user=postgres password=catalina dbname=gobookclub sslmode=disable")
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&User{}, &Book{}, &Role{})
	DB = database
}
