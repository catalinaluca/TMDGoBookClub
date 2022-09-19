package main

import (
	book "gobookclub/controllers/book"
	user "gobookclub/controllers/user"
	"gobookclub/models"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/books", book.FindBooks)
	r.GET("/book/:id", book.FindBook)
	r.POST("/book", book.CreateBook)
	r.PUT("/book/:id", book.UpdateBook)
	r.DELETE("/book/:id", book.DeleteBook)
	r.PUT("/add/waiter", book.AddWaiter)
	r.GET("/users", user.FindUsers)
	r.GET("/user/:id", user.FindUser)
	r.POST("/user", user.CreateUser)
	r.PUT("/user/:id", user.UpdateUser)
	r.DELETE("/user/:id", user.DeleteUser)
	r.PUT("/add/role/user/:id", user.AddRole)
	r.PUT("/borrow/book", book.BorrowBook)
	r.GET("/available/books", book.AvailableBooks)
	r.PUT("/extend/renting", book.ExtendRentingPeriod)
	r.GET("/owned/books/:ownerId", book.OwnedBooks)
	r.PUT("/add/wish", user.AddWish)
	r.Run()
}